package ibapi_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	ibapi "gitlab.com/apisw/go-ibapi"
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
	wrapper := &wrapper{}
	sock := ibapi.NewEClientSocket(wrapper)
	sock.EConnect("locahost", 7496, 1)
	time.Sleep(1000 * time.Millisecond)
	assert.True(t, wrapper.nextId > 0)
	sock.EDisconnect()
	sock.Delete()
}
