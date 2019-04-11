package ibapi_test

import (
	"os"
	"sync"
	"testing"
	"time"

	ibapi "github.com/nu11ptr/go-ibapi"
	"github.com/stretchr/testify/assert"
)

func TestContract(t *testing.T) {
	c := ibapi.NewContract("ES", "FUT", "GLOBEX", "202003", "USD")
	assert.Equal(t, "ES", c.Symbol())
	assert.Equal(t, "FUT", c.SecType())
	assert.Equal(t, "GLOBEX", c.Exchange())
	assert.Equal(t, "202003", c.ContractMonth())
	assert.Equal(t, "USD", c.Currency())
}

func TestLimitOrder(t *testing.T) {
	o := ibapi.NewOrder(123, "BUY", "LMT", 100.0, 56.0, "DAY")
	assert.Equal(t, 123, o.ID())
	assert.Equal(t, "BUY", o.Action())
	assert.Equal(t, "LMT", o.Type())
	assert.Equal(t, 100.0, o.Qty())
	assert.Equal(t, 56.0, o.Price())
	assert.Equal(t, "DAY", o.TIF())
}

func TestStopOrder(t *testing.T) {
	o := ibapi.NewOrder(123, "BUY", "STP", 100.0, 56.0, "DAY")
	assert.Equal(t, 123, o.ID())
	assert.Equal(t, "BUY", o.Action())
	assert.Equal(t, "STP", o.Type())
	assert.Equal(t, 100.0, o.Qty())
	assert.Equal(t, 56.0, o.Price())
	assert.Equal(t, "DAY", o.TIF())
}

func TestMarketOrder(t *testing.T) {
	o := ibapi.NewOrder(123, "BUY", "MKT", 100.0, 56.0, "DAY")
	assert.Equal(t, 123, o.ID())
	assert.Equal(t, "BUY", o.Action())
	assert.Equal(t, "MKT", o.Type())
	assert.Equal(t, 100.0, o.Qty())
	assert.Equal(t, 0.0, o.Price())
	assert.Equal(t, "DAY", o.TIF())
}

type wrapper struct {
	nextId ibapi.OrderID
	lock   sync.Mutex
}

func (w *wrapper) NextValidId(id ibapi.OrderID) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.nextId = id
}

func (w *wrapper) UpdateAccountTime(timeStamp string) {}

func (w *wrapper) Error(id, errorCode int, errorStr string) {}

func (w *wrapper) ConnectionClosed() {}

func (w *wrapper) AccountSummary(reqID int, account, tag, value string) {}

func (w *wrapper) AccountSummaryEnd(reqID int) {}

func TestEClientSocket(t *testing.T) {
	if _, ok := os.LookupEnv("IB_INTEGRATION"); !ok {
		t.Skip("'IB_INTEGRATION' not set - skipping")
	}
	wrapper := &wrapper{nextId: -1}
	client := ibapi.NewIBClient(wrapper)
	assert.True(t, client.Connect("127.0.0.1", 7497, 0))
	go func() {
		for client.IsConnected() {
			client.ProcessMsg()
		}
	}()
	time.Sleep(1000 * time.Millisecond)
	assert.True(t, wrapper.nextId > -1)
	client.Disconnect()
	client.Delete()
}
