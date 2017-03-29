package synth

import (
	"math/rand"
)

func noise(out []float64) {
	for i := range out {
		out[i] = rand.Float64()
	}
}
