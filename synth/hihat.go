package synth

// hihat is a hihat hissing sound. The envelope of which can be controlled. A filter is included to brighten the noise.
type hiHat struct {
	noise  Reader
	env    envelope
	filter filter
	frames hiHatFrames
}

func newHiHat(rise, fall, cutoff float64) *hiHat {
	return &hiHat{
		noise: ReaderFunc(noise),
		env: envelope{
			state: envIdle,
			rise:  rise,
			fall:  fall,
		},
		filter: filter{
			kind:   filterHP,
			poles:  4,
			cutoff: cutoff,
			after:  [4]float64{},
		},
		frames: hiHatFrames{
			noise: NewFrame(),
			env:   NewFrame(),
		},
	}
}

func (h *hiHat) ClockedRead(clock, out []float64) {
	h.noise.Read(h.frames.noise)
	h.env.ClockedRead(clock, h.frames.env)
	for i := range out {
		out[i] = h.filter.tick(h.frames.noise[i] * h.frames.env[i])
	}
}

type hiHatFrames struct {
	noise, env []float64
}
