package engine

import (
	"github.com/brettbuddin/sm808/synth"
	"github.com/gordonklaus/portaudio"
)

// Engine is the interface between the drum machine and PortAudio
type Engine struct {
	stream       *portaudio.Stream
	errors, stop chan error
	reader       synth.Reader
	inFrame      []float64
}

// New returns a new Engine
func New(r synth.Reader) (*Engine, error) {
	if err := portaudio.Initialize(); err != nil {
		return nil, err
	}
	return &Engine{
		errors:  make(chan error),
		stop:    make(chan error),
		reader:  r,
		inFrame: synth.NewFrame(),
	}, nil
}

// Run starts the engine
func (e *Engine) Run() error {
	var err error
	e.stream, err = portaudio.OpenDefaultStream(0, 1, synth.SampleRate, synth.FrameSize, e.callback)
	if err != nil {
		return err
	}

	go func() {
		if err = e.stream.Start(); err != nil {
			e.errors <- err
		}
		<-e.stop

		err = e.stream.Stop()
		if err == nil {
			err = e.stream.Close()
		}
		e.stop <- err
	}()
	return nil
}

// Errors returns a channel that should be consumed to watch for any errors starting the stream
func (e *Engine) Errors() <-chan error {
	return e.errors
}

// Close closes the stream
func (e *Engine) Close() error {
	defer portaudio.Terminate()
	e.stop <- nil
	err := <-e.stop
	close(e.errors)
	close(e.stop)
	return err
}

// callback is the contact point between Go and PortAudio C bindings
func (e *Engine) callback(_, out []float32) {
	e.reader.Read(e.inFrame)
	for i := range out {
		out[i] = float32(e.inFrame[i])
	}
}
