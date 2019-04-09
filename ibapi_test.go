package ibapi_test

import (
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

type wrapper struct {
	nextId ibapi.OrderID
	lock   sync.Mutex
}

func (w *wrapper) NextValidId(id ibapi.OrderID) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.nextId = id
}

func TestEClientSocket(t *testing.T) {
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
