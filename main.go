// sm808 is a drum machine inspired by the Roland TR-808.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gordonklaus/portaudio"

	"github.com/brettbuddin/sm808/engine"
	"github.com/brettbuddin/sm808/synth"
	"github.com/brettbuddin/sm808/tab"
)

func main() {
	if err := command(); err != nil {
		fmt.Println("sm808:", err)
		os.Exit(1)
	}
}

func command() error {
	var (
		tempo   = flag.Int("tempo", 60, "beats per minute")
		steps   = flag.Int("steps", 1, "steps per beat")
		noPrint = flag.Bool("no-print", false, "disable step printing")
	)
	flag.Parse()

	synth.PrintSteps = !*noPrint

	if len(flag.Args()) == 0 {
		return fmt.Errorf("tabfile required")
	}

	t, err := tab.Load(flag.Args()[0])
	if err != nil {
		return err
	}

	voices := []*synth.Voice{}
	for k, v := range t {
		reader, err := synth.GetVoice(k)
		if err != nil {
			return err
		}
		voices = append(voices, synth.NewVoice(k, reader, v))
	}

	m, err := synth.NewMachine(*tempo, *steps, voices...)
	if err != nil {
		return err
	}

	e, err := engine.New(m)
	if err != nil {
		return err
	}
	if err := e.Run(); err != nil {
		return err
	}
	go func() {
		for err := range e.Errors() {
			switch err.(type) {
			case portaudio.Error:
				fmt.Println("engine error:", err)
				os.Exit(1)
			case portaudio.UnanticipatedHostError:
				fmt.Println("engine error:", err)
				os.Exit(1)
			}
		}
	}()

	waitForSignal()
	if err := e.Close(); err != nil {
		return err
	}

	return nil
}

func waitForSignal() {
	sig := make(chan os.Signal)
	done := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		done <- struct{}{}
	}()
	<-done
}
