package tryhard

import (
	"time"
)

type Handler func() error

type TryHard struct {
	store Store

	maximumRetry int
	backOff      int

	defaultKeyTTL int
}

func New(store Store, configs ...func(*TryHard)) *TryHard {
	t := &TryHard{
		store:         store,
		maximumRetry:  3,
		backOff:       5,
		defaultKeyTTL: 300,
	}

	for _, config := range configs {
		config(t)
	}

	return t
}

// Test for create TryHard in test version
func Test(store Store) *TryHard {
	return &TryHard{
		store:         store,
		maximumRetry:  0,
		backOff:       0,
		defaultKeyTTL: 0,
	}
}

func SetMaximumRetry(maximumRetry int) func(*TryHard) {
	return func(t *TryHard) {
		t.maximumRetry = maximumRetry
	}
}

func SetBackOff(backOff int) func(*TryHard) {
	return func(t *TryHard) {
		t.backOff = backOff
	}
}

func SetTimeToLive(second int) func(*TryHard) {
	return func(t *TryHard) {
		t.defaultKeyTTL = second
	}
}

func (t *TryHard) delay(second int) {
	time.Sleep(time.Second * time.Duration(second))
}
