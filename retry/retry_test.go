package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockSleep struct {
	Counter int
}

func (m *mockSleep) Sleep(d time.Duration) {
	m.Counter = m.Counter + 1
}

func TestExecuteSuccessNoRetries(t *testing.T) {
	var task = func() (string, error) {
		return "my result", nil
	}

	mockSleep := &mockSleep{}
	config := Config{
		Times:   5,
		Sleeper: mockSleep,
	}

	result, err := Execute(config, task)

	assert.NoError(t, err)
	assert.Equal(t, "my result", result)
	assert.Equal(t, 0, mockSleep.Counter)
}

func TestExecuteSuccessNoRetriesWithDefaultConstructor(t *testing.T) {
	var task = func() (string, error) {
		return "my result", nil
	}

	result, err := Execute(NewConfig(3), task)

	assert.NoError(t, err)
	assert.Equal(t, "my result", result)
}

func TestExecuteSuccessAfterRetries(t *testing.T) {
	expectedErr := errors.New("Couldn't fetch to the database")
	retryCounter := 0

	var task = func() (string, error) {

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

	assert.NoError(t, err)
	assert.Equal(t, "My result", result)
	assert.Equal(t, 2, mockSleep.Counter)
}

func TestExecuteSuccessAfterRetriesWithCustomExponentialBackoff(t *testing.T) {
	expectedErr := errors.New("Couldn't fetch to the database")
	retryCounter := 0

	var task = func() (string, error) {

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

	assert.NoError(t, err)
	assert.Equal(t, "My result", result)
	assert.Equal(t, 2, mockSleep.Counter)
}

func TestExecuteAlwaysError(t *testing.T) {
	expectedErr := errors.New("Couldn't connect to the database")

	var task = func() (int, error) {
		return 0, expectedErr
	}

	mockSleep := &mockSleep{}

	config := Config{
		Times:   3,
		Sleeper: mockSleep,
	}

	result, err := Execute(config, task)

	assert.Error(t, expectedErr, err)
	assert.Equal(t, 0, result)
	assert.Equal(t, 2, mockSleep.Counter)
}
