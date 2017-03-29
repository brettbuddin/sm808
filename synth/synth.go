// Package synth provides synthesized drum sounds
package synth

import (
	"fmt"
)

// Sampling frequency used by the audio engine
const SampleRate = 44100

var (
	// PrintSteps toggles whether step output should be printed to stdout
	PrintSteps bool

	// Registered voices
	voices = map[string]ClockedReader{
		"closedhihat": newHiHat(ms(1), ms(100), hz(6000)),
		"openhihat":   newHiHat(ms(10), ms(300), hz(5000)),
		"snaredrum": newDrum(drumDef{
			pitch:       hz(250),
			toneRise:    ms(5),
			toneFall:    ms(50),
			noiseRise:   ms(3),
			noiseFall:   ms(200),
			noiseCutoff: hz(5000),
			noiseLevel:  0.5,
			distort:     2,
		}),
		"bassdrum": newDrum(drumDef{
			pitch:       hz(50),
			toneRise:    ms(10),
			toneFall:    ms(1200),
			noiseRise:   ms(5),
			noiseFall:   ms(10),
			noiseCutoff: hz(3000),
			noiseLevel:  1,
			distort:     2,
		}),
	}
)

func GetVoice(name string) (ClockedReader, error) {
	v, ok := voices[name]
	if !ok {
		return nil, fmt.Errorf("no synth voice called %q", name)
	}
	return v, nil
}

func ms(v float64) float64 {
	return v * SampleRate * 0.001
}

func hz(v float64) float64 {
	return v / SampleRate
}

func samplesPerBeat(bpm int) int {
	return int(60 / float64(bpm) * SampleRate)
}

func logf(f string, v ...interface{}) {
	if PrintSteps {
		fmt.Printf(f, v...)
	}
}
