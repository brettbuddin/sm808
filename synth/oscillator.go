package synth

import (
	"math"
)

const twoPi = 2 * math.Pi

// oscillator is a sine wave oscillator
type oscillator struct {
	pitch, phase float64
}

func (o *oscillator) Read(out []float64) {
	for i := range out {
		next := math.Sin(o.phase)
		o.phase += o.pitch * twoPi
		if o.phase <= twoPi {
			o.phase -= twoPi
		}
		out[i] = next
	}
}
