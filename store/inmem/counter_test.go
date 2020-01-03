package inmem

import (
	"testing"
	"time"

	"github.com/devit-tel/gotime"
	"github.com/stretchr/testify/require"
)

func Test_AddAndGetCounter_WithoutTTL(t *testing.T) {
	inmemStore := New()

	inmemStore.AddCounter("x_1")
	xCounter := inmemStore.GetCounter("x_1")
	xExpire, _ := inmemStore.ListTTL["x_1"]
	require.Equal(t, 1, xCounter)
	require.Equal(t, int64(0), xExpire)

	inmemStore.AddCounter("x_1")
	xCounter = inmemStore.GetCounter("x_1")
	xExpire, _ = inmemStore.ListTTL["x_1"]
	require.Equal(t, 2, xCounter)
	require.Equal(t, int64(0), xExpire)

	inmemStore.AddCounter("x_2")
	xCounter = inmemStore.GetCounter("x_2")
	xExpire, _ = inmemStore.ListTTL["x_2"]
	require.Equal(t, 1, xCounter)
	require.Equal(t, int64(0), xExpire)

	inmemStore.AddCounter("x_3")
	xCounter = inmemStore.GetCounter("x_3")
	xExpire, _ = inmemStore.ListTTL["x_3"]
	require.Equal(t, 1, xCounter)
	require.Equal(t, int64(0), xExpire)

	inmemStore.ClearCounter("x_1")
	xCounter = inmemStore.GetCounter("x_1")
	xExpire, _ = inmemStore.ListTTL["x_1"]
	require.Equal(t, 0, xCounter)
	require.Equal(t, int64(0), xExpire)
}

func Test_GetCounter_WithoutTTL_NotFound(t *testing.T) {
	inmemStore := New()

	xCounter := inmemStore.GetCounter("")
	xExpire, _ := inmemStore.ListTTL[""]
	require.Equal(t, 0, xCounter)
	require.Equal(t, int64(0), xExpire)

	xCounter = inmemStore.GetCounter("x_1")
	xExpire, _ = inmemStore.ListTTL["x1"]
	require.Equal(t, 0, xCounter)
	require.Equal(t, int64(0), xExpire)

	xCounter = inmemStore.GetCounter("x_2")
	xExpire, _ = inmemStore.ListTTL["x_2"]
	require.Equal(t, 0, xCounter)
	require.Equal(t, int64(0), xExpire)
}

func Test_GetAndAddCounter_WithTTL(t *testing.T) {
	inmemStore := New()

	now := time.Now()
	gotime.Freeze(now)
	ttl := 60

	inmemStore.AddCounterWithTTL("tx_1", ttl)
	txCounter := inmemStore.GetCounter("tx_1")
	txExpire, _ := inmemStore.ListTTL["tx_1"]
	require.Equal(t, 1, txCounter)
	require.Equal(t, now.Add(time.Duration(ttl)*time.Second).Unix(), txExpire)

	inmemStore.AddCounterWithTTL("tx_1", ttl)
	txCounter = inmemStore.GetCounter("tx_1")
	txExpire, _ = inmemStore.ListTTL["tx_1"]
	require.Equal(t, 2, txCounter)
	require.Equal(t, now.Add(time.Duration(ttl)*time.Second).Unix(), txExpire)
}

func Test_ExpireCounter_WithTTL(t *testing.T) {
	inmemStore := New()

	now := time.Now()
	gotime.Freeze(now)

	inmemStore.AddCounterWithTTL("zz_1", -10)
	inmemStore.AddCounterWithTTL("zz_2", 10)
	time.Sleep(1 * time.Second)
	txCounter := inmemStore.GetCounter("zz_1")
	txExpire, _ := inmemStore.ListTTL["zz_1"]
	require.Equal(t, 0, txCounter)
	require.Equal(t, int64(0), txExpire)

	txCounter = inmemStore.GetCounter("zz_2")
	txExpire, _ = inmemStore.ListTTL["zz_2"]
	require.Equal(t, 1, txCounter)
	require.Equal(t, now.Add(10*time.Second).Unix(), txExpire)
}
