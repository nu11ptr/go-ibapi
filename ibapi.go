package ibapi

//go:generate go tool cgo ibapi.go

/*
#cgo CXXFLAGS: -std=c++11
#cgo CPPFLAGS: -Itws-api
#cgo LDFLAGS: tws-api/libTwsSocketClient.so

#include <stdlib.h>
#include "ibapi.h"
*/
import "C"

import (
	"runtime"
	"sync"
	"unsafe"
)

// *** Contract ***

// Contract represents a contract
type Contract struct {
	contract *C.Contract
}

// NewContract creates a new contract
func NewContract(sym, secType, exch, contractMonth, currency string) *Contract {
	cSym, cSecType, cExch, cContractMonth, cCurrency := C.CString(sym), C.CString(secType),
		C.CString(exch), C.CString(contractMonth), C.CString(currency)
	// NOTE: Since the underlying class uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		C.free(unsafe.Pointer(cSym))
		C.free(unsafe.Pointer(cSecType))
		C.free(unsafe.Pointer(cExch))
		C.free(unsafe.Pointer(cContractMonth))
		C.free(unsafe.Pointer(cCurrency))
	}()
	c := &Contract{contract: C.new_contract(cSym, cSecType, cExch, cContractMonth, cCurrency)}
	runtime.SetFinalizer(c, deleteContract)
	return c
}

func deleteContract(c *Contract) {
	C.delete_contract(c.contract)
}

// Symbol returns the symbol from the contract
func (c *Contract) Symbol() string {
	return C.GoString(C.contract_symbol(c.contract))
}

// SecType returns the security type of the contract
func (c *Contract) SecType() string {
	return C.GoString(C.contract_sec_type(c.contract))
}

// Exchange returns the exchange of the contract
func (c *Contract) Exchange() string {
	return C.GoString(C.contract_exchange(c.contract))
}

// ContractMonth returns the contract month of the futures contract
func (c *Contract) ContractMonth() string {
	return C.GoString(C.contract_month(c.contract))
}

// Currency returns the currency used for the contract
func (c *Contract) Currency() string {
	return C.GoString(C.contract_currency(c.contract))
}

// *** Order ***

// Order represents a broker order ticket
type Order struct {
	order *C.Order
}

// NewOrder creates a new broker order ticket
func NewOrder(orderID int, action, orderType string, qty, price float64, tif string) *Order {
	cAction, cType, cTIF := C.CString(action), C.CString(orderType), C.CString(tif)
	// NOTE: Since the underlying class uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		C.free(unsafe.Pointer(cAction))
		C.free(unsafe.Pointer(cType))
		C.free(unsafe.Pointer(cTIF))
	}()
	o := &Order{order: C.new_order(C.int(orderID), cAction, cType, C.double(qty), C.double(price), cTIF)}
	runtime.SetFinalizer(o, deleteOrder)
	return o
}

func deleteOrder(o *Order) {
	C.delete_order(o.order)
}

// ID returns the order ID
func (o *Order) ID() int {
	return int(C.order_id(o.order))
}

// Action returns the order action (buy, sell, etc.)
func (o *Order) Action() string {
	return C.GoString(C.order_action(o.order))
}

// Type returns the order type (LMT, MkT, STP, etc.)
func (o *Order) Type() string {
	return C.GoString(C.order_type(o.order))
}

// Qty returns the order quantity
func (o *Order) Qty() float64 {
	return float64(C.order_qty(o.order))
}

// Price returns the order price
func (o *Order) Price() float64 {
	return float64(C.order_price(o.order))
}

// TIF returns the order time in force (DAY, GTC, etc.)
func (o *Order) TIF() string {
	return C.GoString(C.order_tif(o.order))
}

// *** EWrapper ***

type wrappers struct {
	m    map[C.long]EWrapper
	next C.long
	lock sync.Mutex
}

var (
	w = wrappers{m: make(map[C.long]EWrapper, 64)}
)

// OrderID represents an IB order ID
type OrderID = C.OrderId

// EWrapper represesnts an interface of IB callbacks
type EWrapper interface {
	NextValidId(orderID OrderID)

	UpdateAccountTime(timeStamp string)

	Error(id, errorCode int, errorStr string)

	ConnectionClosed()

	AccountSummary(reqID int, account, tag, value string)

	AccountSummaryEnd(reqID int)
}

func findEWrapper(id C.long) EWrapper {
	w.lock.Lock()
	defer w.lock.Unlock()
	wrapper, ok := w.m[id]
	if !ok {
		// TODO: log error
		return nil
	}
	return wrapper
}

//export updateAccountTimeCallback
func updateAccountTimeCallback(id C.long, timeStamp *C.char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.UpdateAccountTime(C.GoString(timeStamp))
	}
}

//export errorCallback
func errorCallback(id C.long, errID, errCode C.int, errStr *C.char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.Error(int(errID), int(errCode), C.GoString(errStr))
	}
}

//export connectionClosedCallback
func connectionClosedCallback(id C.long) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.ConnectionClosed()
	}
}

//export nextValidIDCallback
func nextValidIDCallback(id C.long, orderID C.OrderId) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.NextValidId(orderID)
	}
}

//export accountSummaryCallback
func accountSummaryCallback(id C.long, reqID C.int, account, tag, value, currency *C.char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummary(int(reqID), C.GoString(account), C.GoString(tag), C.GoString(currency))
	}
}

//export accountSummaryEndCallback
func accountSummaryEndCallback(id C.long, reqID C.int) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummaryEnd(int(reqID))
	}
}

// *** EClientSocket ***

const timeoutMs = 2000

// IBClient represents an IB client socket
type IBClient struct {
	client *C.IBClient
	id     C.long
}

// NewIBClient returns a new client socket with the given EWrapper callbacks
func NewIBClient(wrapper EWrapper) *IBClient {
	w.lock.Lock()
	next := w.next
	w.m[next] = wrapper
	w.next++
	w.lock.Unlock()

	return &IBClient{client: C.new_client(next, timeoutMs), id: next}
}

// Connect attempts to connect to TWS/IBGateway on the given host/port and client ID
func (c *IBClient) Connect(host string, port, clientID int) bool {
	cHost := C.CString(host)
	// This does NOT use std::string in underlying code, but should be not needed
	// when this call returns
	defer C.free(unsafe.Pointer(cHost))
	return bool(C.client_connect(c.client, cHost, C.int(port), C.int(clientID)))
}

// Disconnect attempts to disconnect from TWS/IBGateway
func (c *IBClient) Disconnect() {
	C.client_disconnect(c.client)
}

// IsConnected returns the connection state of the client
func (c *IBClient) IsConnected() bool {
	return bool(C.client_is_connected(c.client))
}

// ProcessMsg processes the next msg waiting on the client
func (c *IBClient) ProcessMsg() {
	C.client_process_msg(c.client)
}

// Delete frees the underlying CPP resources and removes the wrapper from the map
func (c *IBClient) Delete() {
	// First, get rid of the underlying socket to prevent callbacks
	C.delete_client(c.client)
	// Now remove the reference to the ewrapper from the map
	w.lock.Lock()
	delete(w.m, c.id)
	w.lock.Unlock()
}

// ReqAccountSummary requests the summaries for all accounts
func (c *IBClient) ReqAccountSummary(reqID int, group, tags string) {
	cGroup, cTags := C.CString(group), C.CString(tags)
	// NOTE: Since the underlying code uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		C.free(unsafe.Pointer(cGroup))
		C.free(unsafe.Pointer(cTags))
	}()
	C.client_req_account_summ(c.client, C.int(reqID), cGroup, cTags)
}

// CancelAccountSummary cancels the account summary info
func (c *IBClient) CancelAccountSummary(reqID int) {
	C.client_cancel_account_summ(c.client, C.int(reqID))
}

// PlaceOrder places an order with the given contracts
func (c *IBClient) PlaceOrder(orderID OrderID, contract *Contract, order *Order) {
	C.client_place_order(c.client, orderID, contract.contract, order.order)
}

// CancelOrder cancels the order with the given order id
func (c *IBClient) CancelOrder(orderID OrderID) {
	C.client_cancel_order(c.client, orderID)
}
