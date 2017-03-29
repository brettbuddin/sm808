package synth

import (
	"math"
)

// envelope is a simple Attack-Decay envelope. Acts as a simple state-machine that moves between idle -> rise -> fall
// and back again. Retriggering is possible while in the fall state.
type envelope struct {
	trigger, lastTrigger, velocity float64
	rise, fall, out                float64
	state                          envStateFn
}

func (e *envelope) ClockedRead(clock, out []float64) {
	for i := range out {
		if clock[i] > 0 {
			e.velocity = clock[i]
		}
		e.trigger = clock[i]
		e.state = e.state(e)
		e.lastTrigger = e.trigger
		out[i] = e.out * e.velocity
	}
}

func (e *envelope) isTriggered() bool {
	return e.lastTrigger == 0 && e.trigger > 0
}

type envStateFn func(*envelope) envStateFn

func envIdle(e *envelope) envStateFn {
	e.out = 0
	if e.isTriggered() {
		return envRise
	}
	return envIdle
}

// Rise uses a logorithmic curve
func envRise(e *envelope) envStateFn {
	multiplier := curve(e.rise)
	base := (1.0 + expRatio) * (1.0 - multiplier)
	e.out = base + e.out*multiplier

	if e.out >= 1 {
		e.out = 1
		return envFall
	}
	return envRise
}

// Fall uses a logorithmic curve
func envFall(e *envelope) envStateFn {
	multiplier := curve(e.fall)
	base := -expRatio * (1.0 - multiplier)
	e.out = base + e.out*multiplier

	if e.isTriggered() {
		return envRise
	}

	if e.out <= 0 {
		e.out = 0
		return envIdle
	}
	return envFall
}

// Controls the amount of inflection at the edges of the curve
const expRatio = 0.01

// Controlled exponential curve
func curve(speed float64) float64 {
	r := (1 + expRatio) / expRatio
	return math.Exp(-math.Log(r) / speed)
}
