package retry

import (
	"time"

	"github.com/Talento90/goliath/sleep"
)

// Config retry mechanism
type Config struct {
	//Number of retries to be applied
	Times int
	//ExponentialBackoff function that calculates the retry delay
	ExponentialBackoff func(retryCount int) time.Duration
	//Sleeper pauses the execution of the current go routine for x milliseconds
	Sleeper sleep.Sleeper
}

func NewConfig(retryCount int) Config {
	return Config{
		Sleeper:            sleep.New(),
		Times:              retryCount,
		ExponentialBackoff: defaultExponentialBackoff,
	}
}

func defaultExponentialBackoff(retryCount int) time.Duration {
	return time.Duration(retryCount * 100)
}

// Execute the task and retries when the task returns an error
func Execute[T any](config Config, task func() (T, error)) (T, error) {
	var retryErr error
	var defaultResult T

	for i := 0; i < config.Times; i++ {
		result, err := task()

		if err == nil {
			return result, nil
		}

		if i == config.Times-1 {
			retryErr = err
			break
		}

		var delay time.Duration

		if config.ExponentialBackoff != nil {
			delay = config.ExponentialBackoff(i + 1)
		} else {
			delay = defaultExponentialBackoff(i + 1)
		}

		config.Sleeper.Sleep(delay * time.Millisecond)
	}

	return defaultResult, retryErr
}
