package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUtcClock(t *testing.T) {
	c := NewUtcClock()
	now := c.Now().Format(time.RFC822)

	assert.Contains(t, now, "UTC")
}

func TestNewNoLocation(t *testing.T) {
	c := New(nil)
	now := c.Now().Format(time.RFC822)

	assert.Contains(t, now, "UTC")
}

func TestNewWithLocation(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Error(err)
	}

	c := New(loc)
	now := c.Now().Format(time.RFC822)

	assert.Contains(t, now, "CST")
}
