package synth

import (
	"testing"

	"github.com/brettbuddin/sm808/test/assert"
)

func TestMachineStepAdvancment(t *testing.T) {
	// (60 / 100000bpm) * 44100hz = 264.6
	// We should expect a clock tick every 26 samples.

	bpm := 100000
	stepsPerBeat := 10

	expected := NewFrame()
	expected[26] = 0.4

	reader := &mockReader{
		fn: func(clock, out []float64) {
			var i int
			for j := 0; j < FrameSize; j += 26 {
				if i%2 == 0 {
					assert.Equal(t, clock[j], 1.0)
				} else {
					assert.Equal(t, clock[j], 0.4)
				}
				i++
			}
		},
	}
	v := NewVoice("test", reader, []rune{'X', 'x'})

	m, err := NewMachine(bpm, stepsPerBeat, v)
	assert.Equal(t, err, nil)

	m.Read(NewFrame())
	assert.Equal(t, reader.called, true)
}
