package sleep

import (
	"testing"
)

func TestSleep(t *testing.T) {
	sleeper := New()

	sleeper.Sleep(0)
}
