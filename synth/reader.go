package synth

// Size of the outgoing audio buffer
const FrameSize = 512

// ClockedReader is a Reader that also needs to be aware of a clock or series of triggers.
type ClockedReader interface {
	ClockedRead(clock, out []float64)
}

// Reader is something that can read audio data into a buffer
type Reader interface {
	Read([]float64)
}

// ReaderFunc is a function that implements Reader
type ReaderFunc func(out []float64)

func (f ReaderFunc) Read(out []float64) {
	f(out)
}

// NewFrame creates a new frame the size of the outgoing audio buffer
func NewFrame() []float64 {
	return make([]float64, FrameSize)
}
