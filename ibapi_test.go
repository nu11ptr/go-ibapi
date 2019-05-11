package ibapi_test

import (
	"fmt"
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

type account struct {
	account string
	keyVal  map[string]string
}

func (a *account) AddKeyVal(key, val string) {
	a.keyVal[key] = val
}

type wrapper struct {
	nextId   ibapi.OrderID
	lock     sync.Mutex
	accounts map[string]*account
	acctDone bool
}

func (w *wrapper) NextValidId(id ibapi.OrderID) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.nextId = id
}

func (w *wrapper) UpdateAccountTime(timeStamp string) {
	fmt.Printf("Account time: %s\n", timeStamp)
}

func (w *wrapper) Error(id, errorCode int, errorStr string) {
	fmt.Printf("Error ID:%d Code:%d Msg:%s\n", id, errorCode, errorStr)
}

func (w *wrapper) ConnectionClosed() {
	fmt.Println("Closed")
}

func (w *wrapper) AccountSummary(reqID int, acctID, tag, value string) {
	w.lock.Lock()
	defer w.lock.Unlock()

	acct, ok := w.accounts[acctID]
	if !ok {
		acct = &account{account: acctID, keyVal: make(map[string]string, 20)}
		w.accounts[acctID] = acct
	}
	acct.AddKeyVal(tag, value)
}

func (w *wrapper) AccountSummaryEnd(reqID int) {
	w.lock.Lock()
	defer w.lock.Unlock()

	w.acctDone = true
}

func (w *wrapper) OpenOrder(orderID ibapi.OrderID, contract *ibapi.Contract, order *ibapi.Order) {}

func (w *wrapper) OrderStatus(orderID ibapi.OrderID, status string, filled, remaining, avgFillPrice float64,
	permID, parentID int, lastFillPrice float64, clientID int, whyHeld string, mktCapPrice float64) {
}

func (w *wrapper) OrderBound(orderID ibapi.OrderID, apiClientID, apiOrderID int) {}

func (w *wrapper) OpenOrderEnd() {}

func (w *wrapper) ExecDetails(reqID int, contract *ibapi.Contract, exec *ibapi.Execution) {}

func (w *wrapper) ExecDetailsEnd(reqID int) {}

func TestEClientSocket(t *testing.T) {
	if _, ok := os.LookupEnv("IB_INTEGRATION"); !ok {
		t.Skip("'IB_INTEGRATION' not set - skipping")
	}
	wrapper := &wrapper{nextId: -1, accounts: make(map[string]*account, 5)}
	client := ibapi.NewIBClient(wrapper)

	t.Run("Connect", func(t *testing.T) {
		assert.True(t, client.Connect("127.0.0.1", 7497, 0))
		go func() {
			for client.IsConnected() {
				client.ProcessMsg()
			}
		}()
		time.Sleep(1000 * time.Millisecond)
		assert.True(t, wrapper.nextId > -1)
	})
	t.Run("AccountSummary", func(t *testing.T) {
		client.ReqAccountSummary(1, "All", "AccountType,NetLiquidation")
		time.Sleep(1000 * time.Millisecond)
		assert.NotEmpty(t, wrapper.accounts)
		for _, acct := range wrapper.accounts {
			assert.NotEmpty(t, acct.account)
			assert.Len(t, acct.keyVal, 2)
		}
		assert.True(t, wrapper.acctDone)
	})

	client.Disconnect()
	time.Sleep(1000 * time.Millisecond)
	client.Delete()
}
