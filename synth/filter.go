package synth

import "math"

const (
	filterLP int = iota
	filterHP
)

// filter is a N-pole low-pass or high-pass filter
type filter struct {
	kind, poles int
	cutoff      float64
	after       [4]float64
}

func (f *filter) tick(v float64) float64 {
	cutoff := twoPi * math.Abs(f.cutoff)

	last := v
	for i := 0; i < f.poles; i++ {
		f.after[i] += (-f.after[i] + last) * cutoff
		last = f.after[i]
	}

	switch f.kind {
	case filterHP:
		return v - f.after[3]
	}
	return f.after[3]
}
