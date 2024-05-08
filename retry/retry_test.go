package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockSleep struct {
	Counter int
}

func (m *mockSleep) Sleep(time.Duration) {
	m.Counter = m.Counter + 1
}

func TestExecuteSuccessNoRetries(t *testing.T) {
	task := func() (string, error) {
		return "my result", nil
	}

	mockSleep := &mockSleep{}
	config := Config{
		Times:   5,
		Sleeper: mockSleep,
	}

	result, err := Execute(config, task)

	require.NoError(t, err)
	require.Equal(t, "my result", result)
	require.Equal(t, 0, mockSleep.Counter)
}

func TestExecuteSuccessNoRetriesWithDefaultConstructor(t *testing.T) {
	task := func() (string, error) {
		return "my result", nil
	}

	result, err := Execute(NewConfig(3), task)

	require.NoError(t, err)
	require.Equal(t, "my result", result)
}

func TestExecuteSuccessAfterRetries(t *testing.T) {
	expectedErr := errors.New("Couldn't fetch to the database")
	retryCounter := 0

	task := func() (string, error) {
		if retryCounter == 2 {
			return "My result", nil
		}

		retryCounter++

		return "", expectedErr
	}

	mockSleep := &mockSleep{}

	config := Config{
		Times:   5,
		Sleeper: mockSleep,
	}

	result, err := Execute(config, task)

	require.NoError(t, err)
	require.Equal(t, "My result", result)
	require.Equal(t, 2, mockSleep.Counter)
}

func TestExecuteSuccessAfterRetriesWithCustomExponentialBackoff(t *testing.T) {
	expectedErr := errors.New("Couldn't fetch to the database")
	retryCounter := 0

	task := func() (string, error) {
		if retryCounter == 2 {
			return "My result", nil
		}

		retryCounter++

		return "", expectedErr
	}

	mockSleep := &mockSleep{}

	config := Config{
		Times:   5,
		Sleeper: mockSleep,
		ExponentialBackoff: func(retryCount int) time.Duration {
			return time.Duration(retryCount)
		},
	}

	result, err := Execute(config, task)

	require.NoError(t, err)
	require.Equal(t, "My result", result)
	require.Equal(t, 2, mockSleep.Counter)
}

func TestExecuteAlwaysError(t *testing.T) {
	expectedErr := errors.New("Couldn't connect to the database")

	task := func() (int, error) {
		return 0, expectedErr
	}

	mockSleep := &mockSleep{}

	config := Config{
		Times:   3,
		Sleeper: mockSleep,
	}

	result, err := Execute(config, task)

	require.ErrorIs(t, expectedErr, err)
	require.Equal(t, 0, result)
	require.Equal(t, 2, mockSleep.Counter)
}
