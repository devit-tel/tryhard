package tryhard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/devit-tel/tryhard/internal/mocks"
)

func TestTryHard_Do_Success(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(5), SetBackOff(0))

	var workCounter int
	try.Do(func() error {
		workCounter++
		return nil
	})

	require.Equal(t, 1, workCounter)
	mockStore.AssertExpectations(t)
}

func TestTryHard_Do_MaximumRetry(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(5), SetBackOff(0))

	var workCounter int
	try.Do(func() error {
		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 5, workCounter)
	mockStore.AssertExpectations(t)
}

func TestTryHard_Do_FailedAndSuccess(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(5), SetBackOff(0))

	var workCounter int
	try.Do(func() error {
		if workCounter == 3 {
			return nil
		}

		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 3, workCounter)
	mockStore.AssertExpectations(t)
}

func TestTryHard_DoWithKey_Success(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(3), SetBackOff(0))

	mockStore.On("GetCounter", "key_1").Times(1).Return(0)
	mockStore.On("ClearCounter", "key_1")

	var workCounter int
	try.DoWithKey("key_1", func() error {
		workCounter++
		return nil
	})

	require.Equal(t, 1, workCounter)
	mockStore.AssertExpectations(t)
}

func TestTryHard_DoWithKey_Failed_MaximumRetry(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(2), SetBackOff(0))

	mockStore.On("AddCounterWithTTL", "key_1", 0)
	mockStore.On("GetCounter", "key_1").Return(0).Times(1)
	mockStore.On("GetCounter", "key_1").Return(1).Times(1)
	mockStore.On("GetCounter", "key_1").Return(2).Times(1)
	mockStore.On("GetCounter", "key_1").Return(3).Times(1)

	var workCounter int
	try.DoWithKey("key_1", func() error {
		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 3, workCounter)
	mockStore.AssertExpectations(t)
}

func TestTryHard_DoWithKey_Exceed_MaximumRetry(t *testing.T) {
	mockStore := &mocks.Store{}
	try := New(mockStore, SetMaximumRetry(2), SetBackOff(0))

	mockStore.On("GetCounter", "key_1").Return(3).Times(1)

	var workCounter int
	try.DoWithKey("key_1", func() error {
		workCounter++
		return errors.New("unknown error")
	})

	require.Equal(t, 0, workCounter)
	mockStore.AssertExpectations(t)
}
