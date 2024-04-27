package clock

import "time"

// Clock interface allows to get the current time
type Clock interface {
	Now() time.Time
}

type clock struct {
	location *time.Location
}

// NewUtcClock returns a new instace of Clock using UTC location
func NewUtcClock() Clock {
	return &clock{location: time.UTC}
}

// New returns a new instace of Clock.
// If location is not passed then uses UTC.
func New(location *time.Location) Clock {
	c := &clock{}

	if location == nil {
		c.location = time.UTC
	} else {
		c.location = location
	}

	return c
}

func (c clock) Now() time.Time {
	return time.Now().In(c.location)
}
