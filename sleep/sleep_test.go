package sleep

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSleep(t *testing.T) {
	sleeper := New()

	sleeper.Sleep(0)

	require.NotNil(t, sleeper)
}
