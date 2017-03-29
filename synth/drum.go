package synth

import "math"

// drum is a drum that can model many different kinds of drums. Two envelopes are provided to decouple the shaping the
// tone and noise (white) of the drum. A filter is also included to dampen the brightness of the noise. Slight tone
// distortion is possible by the inclusion of tanh function if the gain for the distortion is higher than 0.
type drum struct {
	oscillator
	filter
	noise               Reader
	noiseLevel, distort float64
	envelopes           drumEnvelopes
	frames              drumFrames
}

type drumDef struct {
	pitch                                         float64
	toneRise, toneFall, distort                   float64
	noiseRise, noiseFall, noiseCutoff, noiseLevel float64
}

func newDrum(def drumDef) *drum {
	return &drum{
		noise:      ReaderFunc(noise),
		oscillator: oscillator{pitch: def.pitch},
		noiseLevel: def.noiseLevel,
		distort:    def.distort,
		filter: filter{
			kind:   filterLP,
			poles:  4,
			cutoff: def.noiseCutoff,
			after:  [4]float64{},
		},
		envelopes: drumEnvelopes{
			tone: envelope{
				state: envIdle,
				rise:  def.toneRise,
				fall:  def.toneFall,
			},
			noise: envelope{
				state: envIdle,
				rise:  def.noiseRise,
				fall:  def.noiseFall,
			},
		},
		frames: drumFrames{
			osc:      NewFrame(),
			noise:    NewFrame(),
			noiseEnv: NewFrame(),
			toneEnv:  NewFrame(),
		},
	}
}

func (d *drum) ClockedRead(clock, out []float64) {
	d.oscillator.Read(d.frames.osc)
	d.noise.Read(d.frames.noise)
	d.envelopes.tone.ClockedRead(clock, d.frames.toneEnv)
	d.envelopes.noise.ClockedRead(clock, d.frames.noiseEnv)

	for i := range out {
		tone := d.frames.osc[i] * d.frames.toneEnv[i]
		if d.distort > 0 {
			tone = math.Tanh(tone * d.distort)
		}
		noise := d.noiseLevel * d.filter.tick((d.frames.noise[i] * d.frames.noiseEnv[i]))
		out[i] = tone + noise
	}
}

type drumEnvelopes struct {
	tone, noise envelope
}

type drumFrames struct {
	osc, noise, toneEnv, noiseEnv []float64
}
