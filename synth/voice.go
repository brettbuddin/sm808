package synth

const (
	velocityLow  = 'x'
	velocityHigh = 'X'

	ampLow  = 0.4
	ampHigh = 1
)

// Voice plays a sequence using a synthesized drum sound. The clock signal sent to the underlying Voice contains
// velocity information for dynamics.
type Voice struct {
	name         string
	reader       ClockedReader
	steps        []rune
	step         int
	triggerFrame []float64
	lastClock    float64
}

func NewVoice(name string, reader ClockedReader, steps []rune) *Voice {
	return &Voice{
		name:         name,
		reader:       reader,
		steps:        steps,
		triggerFrame: make([]float64, FrameSize),
	}
}

func (s *Voice) ClockedRead(clock, out []float64) {
	for i := range out {
		if s.lastClock == 0 && clock[i] == 1 {
			switch s.steps[s.step] {
			case velocityLow:
				s.triggerFrame[i] = ampLow
				logf("%s (soft)\n", s.name)
			case velocityHigh:
				s.triggerFrame[i] = ampHigh
				logf("%s (hard)\n", s.name)
			default:
				s.triggerFrame[i] = 0
			}
			s.step = (s.step + 1) % len(s.steps)
		} else {
			s.triggerFrame[i] = 0
		}
		s.lastClock = clock[i]
	}
	s.reader.ClockedRead(s.triggerFrame, out)
}
