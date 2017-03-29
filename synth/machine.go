package synth

// Machine is a drum machine inspired by the Roland TR-808. It supports 3 voices: bassdrum, snaredrum and hihat. These
// voices are sequenced by providing tab.Tab sequences.
type Machine struct {
	samplesPerBeat int
	stepsPerBeat   int
	voices         []*Voice
	voiceFrames    [][]float64
	clockFrame     []float64
	tick           int
}

// NewMachine returns a new Machine set to a particular tempo, time signature and list of sequenced voices.
func NewMachine(bpm, stepsPerBeat int, voices ...*Voice) (*Machine, error) {
	m := &Machine{
		samplesPerBeat: samplesPerBeat(bpm),
		stepsPerBeat:   stepsPerBeat,
		voices:         voices,
		clockFrame:     NewFrame(),
		voiceFrames:    make([][]float64, len(voices)),
	}
	for i := range m.voiceFrames {
		m.voiceFrames[i] = NewFrame()
	}
	return m, nil
}

func (m *Machine) Read(out []float64) {
	// Read clock and voices
	for i := range out {
		if m.tick == 0 {
			logf("-----\n")
			m.clockFrame[i] = 1
		} else {
			m.clockFrame[i] = 0
		}
		m.tick = (m.tick + 1) % (m.samplesPerBeat / m.stepsPerBeat)
	}
	for i, s := range m.voices {
		s.ClockedRead(m.clockFrame, m.voiceFrames[i])
	}

	// Mix voices
	for i := range out {
		var mix float64
		for j := range m.voices {
			mix += m.voiceFrames[j][i]
		}
		out[i] = mix
	}
}
