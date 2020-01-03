package tryhard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/devit-tel/tryhard/store/inmem"
)

func TestTryHard_Inmem_Do_Success(t *testing.T) {
	try := New(inmem.New())

	var workCounter int
	try.Do(func() error {
		workCounter++
		return nil
	})

	require.Equal(t, 1, workCounter)
}

func TestTryHard_Inmem_Do_MaximumRetry(t *testing.T) {
	try := New(inmem.New(), SetMaximumRetry(2), SetBackOff(0))

	var workCounter int
	try.Do(func() error {
		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 2, workCounter)
}

func TestTryHard_Inmem_Do_FailedAndSuccess(t *testing.T) {
	try := New(inmem.New(), SetMaximumRetry(5), SetBackOff(0))

	var workCounter int
	try.Do(func() error {
		if workCounter == 3 {
			return nil
		}

		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 3, workCounter)
}

func TestTryHard_Inmem_DoWithKey_Success(t *testing.T) {
	try := New(inmem.New(), SetMaximumRetry(5), SetBackOff(0))

	var doNotCount, workCounter int
	try.DoWithKey("key_1", func() error {
		workCounter++
		return errors.New("unknown error")
	})

	try.DoWithKey("key_1", func() error {
		doNotCount++
		return nil
	})

	require.Equal(t, 6, workCounter)
	require.Equal(t, 0, doNotCount)
}
