package sleep

import "time"

// Sleeper pauses the current go routine
type Sleeper interface {
	Sleep(time.Duration)
}

type sleeper struct{}

// New returns a new sleep
func New() Sleeper {
	return sleeper{}
}

// Sleep pauses the execution of the current go routine
func (sleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}
