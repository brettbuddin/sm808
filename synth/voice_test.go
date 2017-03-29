package synth

import (
	"testing"

	"github.com/brettbuddin/sm808/test/assert"
)

type mockReader struct {
	fn     func(clock, out []float64)
	called bool
}

func (r *mockReader) ClockedRead(clock, out []float64) {
	r.called = true
	r.fn(clock, out)
}

func TestVoiceClockedRead(t *testing.T) {
	clock := NewFrame()
	clock[0] = 1
	clock[1] = 0
	clock[2] = 1
	clock[3] = 0
	clock[4] = 1

	frame := make([]float64, len(clock))
	pattern := []rune{'X', '_', 'x'}

	expected := NewFrame()
	expected[0] = 1
	expected[1] = 0
	expected[2] = 0
	expected[3] = 0
	expected[4] = 0.4

	reader := &mockReader{
		fn: func(clock, out []float64) {
			assert.Equal(t, clock, expected)
		},
	}

	v := NewVoice("test", reader, pattern)
	v.ClockedRead(clock, frame)
	assert.Equal(t, reader.called, true)
}
