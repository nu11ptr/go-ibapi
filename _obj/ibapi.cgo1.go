// Code generated by cmd/cgo; DO NOT EDIT.

//line ibapi.go:1:1
package ibapi; import _cgo_unsafe "unsafe"

//go:generate go tool cgo ibapi.go

/*
#cgo CXXFLAGS: -std=c++11
#cgo CPPFLAGS: -Itws-api
#cgo LDFLAGS: tws-api/libTwsSocketClient.so

#include <stdlib.h>
#include "ibapi.h"
*/
import _ "unsafe"

import (
	"runtime"
	"sync"
	"unsafe"
)

// *** Contract ***

// Contract represents a contract
type Contract struct {
	contract *_Ctype_struct_Contract
}

// NewContract creates a new contract
func NewContract(sym, secType, exch, contractMonth, currency string) *Contract {
	cSym, cSecType, cExch, cContractMonth, cCurrency := (_Cfunc_CString)(sym), (_Cfunc_CString)(secType),
		(_Cfunc_CString)(exch), (_Cfunc_CString)(contractMonth), (_Cfunc_CString)(currency)
	// NOTE: Since the underlying class uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cSym))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cSecType))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cExch))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cContractMonth))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cCurrency))
	}()
	c := &Contract{contract: (_Cfunc_new_contract)(cSym, cSecType, cExch, cContractMonth, cCurrency)}
	runtime.SetFinalizer(c, deleteContract)
	return c
}

func deleteContract(c *Contract) {
	func(_cgo0 *_Ctype_struct_Contract) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_delete_contract)(_cgo0);}(c.contract)
}

// Symbol returns the symbol from the contract
func (c *Contract) Symbol() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Contract) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_contract_symbol)(_cgo0);}(c.contract))
}

// SecType returns the security type of the contract
func (c *Contract) SecType() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Contract) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_contract_sec_type)(_cgo0);}(c.contract))
}

// Exchange returns the exchange of the contract
func (c *Contract) Exchange() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Contract) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_contract_exchange)(_cgo0);}(c.contract))
}

// ContractMonth returns the contract month of the futures contract
func (c *Contract) ContractMonth() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Contract) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_contract_month)(_cgo0);}(c.contract))
}

// Currency returns the currency used for the contract
func (c *Contract) Currency() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Contract) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_contract_currency)(_cgo0);}(c.contract))
}

// *** Order ***

// Order represents a broker order ticket
type Order struct {
	order *_Ctype_struct_Order
}

// NewOrder creates a new broker order ticket
func NewOrder(orderID int, action, orderType string, qty, price float64, tif string) *Order {
	cAction, cType, cTIF := (_Cfunc_CString)(action), (_Cfunc_CString)(orderType), (_Cfunc_CString)(tif)
	// NOTE: Since the underlying class uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cAction))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cType))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cTIF))
	}()
	o := &Order{order: (_Cfunc_new_order)(_Ctype_int(orderID), cAction, cType, _Ctype_double(qty), _Ctype_double(price), cTIF)}
	runtime.SetFinalizer(o, deleteOrder)
	return o
}

func deleteOrder(o *Order) {
	func(_cgo0 *_Ctype_struct_Order) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_delete_order)(_cgo0);}(o.order)
}

// ID returns the order ID
func (o *Order) ID() int {
	return int(func(_cgo0 *_Ctype_struct_Order) _Ctype_int {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_id)(_cgo0);}(o.order))
}

// Action returns the order action (buy, sell, etc.)
func (o *Order) Action() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Order) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_action)(_cgo0);}(o.order))
}

// Type returns the order type (LMT, MkT, STP, etc.)
func (o *Order) Type() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Order) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_type)(_cgo0);}(o.order))
}

// Qty returns the order quantity
func (o *Order) Qty() float64 {
	return float64(func(_cgo0 *_Ctype_struct_Order) _Ctype_double {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_qty)(_cgo0);}(o.order))
}

// Price returns the order price
func (o *Order) Price() float64 {
	return float64(func(_cgo0 *_Ctype_struct_Order) _Ctype_double {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_price)(_cgo0);}(o.order))
}

// TIF returns the order time in force (DAY, GTC, etc.)
func (o *Order) TIF() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Order) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_order_tif)(_cgo0);}(o.order))
}

// *** ExecutionFilter ***

// ExecutionFilter represents an execution filter
type ExecutionFilter struct {
	filter *_Ctype_struct_ExecutionFilter
}

// NewFilterParams creates a new execution filter with parameters
func NewFilterParams(clientID int, acctCode, time, symbol, secType, exchange, side string) *ExecutionFilter {
	cAcctCode, cTime, cSymbol, cSecType, cExchange, cSide := (_Cfunc_CString)(acctCode), (_Cfunc_CString)(time),
		(_Cfunc_CString)(symbol), (_Cfunc_CString)(secType), (_Cfunc_CString)(exchange), (_Cfunc_CString)(side)
	// NOTE: Since the underlying class uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cAcctCode))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cTime))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cSymbol))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cSecType))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cExchange))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cSide))
	}()
	f := &ExecutionFilter{filter: (_Cfunc_new_exec_filter)(_Ctype_long(clientID), cAcctCode, cTime, cSymbol,
		cSecType, cExchange, cSide)}
	runtime.SetFinalizer(f, deleteFilter)
	return f
}

// NewFilter creates a new blank filter
func NewFilter() *ExecutionFilter {
	return NewFilterParams(0, "", "", "", "", "", "")
}

func deleteFilter(f *ExecutionFilter) {
	func(_cgo0 *_Ctype_struct_ExecutionFilter) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_delete_exec_filter)(_cgo0);}(f.filter)
}

// *** Execution ***

// Execution represents an order execution
type Execution struct {
	exec *_Ctype_struct_Execution
}

// ID returns the execution ID
func (e *Execution) ID() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Execution) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_id)(_cgo0);}(e.exec))
}

// Time returns the time of the execution
func (e *Execution) Time() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Execution) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_time)(_cgo0);}(e.exec))
}

// AccountNum returns the account number in the execution
func (e *Execution) AccountNum() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Execution) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_account_num)(_cgo0);}(e.exec))
}

// Exchange returns the exchange the execution occurred on
func (e *Execution) Exchange() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Execution) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_exchange)(_cgo0);}(e.exec))
}

// Side returns whether this was a buy/sell/short/cover/etc.
func (e *Execution) Side() string {
	return (_Cfunc_GoString)(func(_cgo0 *_Ctype_struct_Execution) *_Ctype_char {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_side)(_cgo0);}(e.exec))
}

// Shares returns the number of shares executed
func (e *Execution) Shares() float64 {
	return float64(func(_cgo0 *_Ctype_struct_Execution) _Ctype_double {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_shares)(_cgo0);}(e.exec))
}

// Price returns the price that executed
func (e *Execution) Price() float64 {
	return float64(func(_cgo0 *_Ctype_struct_Execution) _Ctype_double {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_price)(_cgo0);}(e.exec))
}

// AvgPrice returns the average price of the execution
func (e *Execution) AvgPrice() float64 {
	return float64(func(_cgo0 *_Ctype_struct_Execution) _Ctype_double {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_avg_price)(_cgo0);}(e.exec))
}

// OrderID returns the associated orderID for this execution
func (e *Execution) OrderID() OrderID {
	return OrderID(func(_cgo0 *_Ctype_struct_Execution) _Ctype_long {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_exec_order_id)(_cgo0);}(e.exec))
}

// *** EWrapper ***

type wrappers struct {
	m    map[_Ctype_long]EWrapper
	next _Ctype_long
	lock sync.Mutex
}

var (
	w = wrappers{m: make(map[_Ctype_long]EWrapper, 64)}
)

// OrderID represents an IB order ID
type OrderID = _Ctype_OrderId

// EWrapper represesnts an interface of IB callbacks
type EWrapper interface {
	NextValidId(orderID OrderID)

	UpdateAccountTime(timeStamp string)

	Error(id, errorCode int, errorStr string)

	ConnectionClosed()

	AccountSummary(reqID int, account, tag, value string)

	AccountSummaryEnd(reqID int)

	OpenOrder(orderID OrderID, contract *Contract, order *Order)

	OrderStatus(orderID OrderID, status string, filled, remaining, avgFillPrice float64,
		permID, parentID int, lastFillPrice float64, clientID int, whyHeld string, mktCapPrice float64)

	OrderBound(orderID OrderID, apiClientID, apiOrderID int)

	OpenOrderEnd()

	ExecDetails(reqID int, contract *Contract, exec *Execution)

	ExecDetailsEnd(reqID int)
}

func findEWrapper(id _Ctype_long) EWrapper {
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
func updateAccountTimeCallback(id _Ctype_long, timeStamp *_Ctype_char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.UpdateAccountTime((_Cfunc_GoString)(timeStamp))
	}
}

//export errorCallback
func errorCallback(id _Ctype_long, errID, errCode _Ctype_int, errStr *_Ctype_char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.Error(int(errID), int(errCode), (_Cfunc_GoString)(errStr))
	}
}

//export connectionClosedCallback
func connectionClosedCallback(id _Ctype_long) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.ConnectionClosed()
	}
}

//export nextValidIDCallback
func nextValidIDCallback(id _Ctype_long, orderID _Ctype_OrderId) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.NextValidId(orderID)
	}
}

//export accountSummaryCallback
func accountSummaryCallback(id _Ctype_long, reqID _Ctype_int, account, tag, value, currency *_Ctype_char) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummary(int(reqID), (_Cfunc_GoString)(account), (_Cfunc_GoString)(tag), (_Cfunc_GoString)(currency))
	}
}

//export accountSummaryEndCallback
func accountSummaryEndCallback(id _Ctype_long, reqID _Ctype_int) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummaryEnd(int(reqID))
	}
}

//export openOrderCallback
func openOrderCallback(id _Ctype_long, orderID OrderID, contract *_Ctype_struct_Contract, order *_Ctype_struct_Order) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.OpenOrder(orderID, &Contract{contract: contract}, &Order{order: order})
	}
}

//export orderStatusCallback
func orderStatusCallback(id _Ctype_long, orderID OrderID, status *_Ctype_char, filled, remaining,
	avgFillPrice _Ctype_double, permID, parentID _Ctype_int, lastFillPrice _Ctype_double, clientID _Ctype_int,
	whyHeld *_Ctype_char, mktCapPrice _Ctype_double) {

	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.OrderStatus(orderID, (_Cfunc_GoString)(status), float64(filled), float64(remaining),
			float64(avgFillPrice), int(permID), int(parentID), float64(lastFillPrice), int(clientID),
			(_Cfunc_GoString)(whyHeld), float64(mktCapPrice))
	}
}

//export orderBoundCallback
func orderBoundCallback(id _Ctype_long, orderID _Ctype_long, apiClientID, apiOrderID _Ctype_int) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.OrderBound(OrderID(orderID), int(apiClientID), int(apiOrderID))
	}
}

//export openOrderEndCallback
func openOrderEndCallback(id _Ctype_long) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.OpenOrderEnd()
	}
}

//export execDetailsCallback
func execDetailsCallback(id _Ctype_long, reqID _Ctype_int, contract *_Ctype_struct_Contract, exec *_Ctype_struct_Execution) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.ExecDetails(int(reqID), &Contract{contract: contract}, &Execution{exec: exec})
	}
}

//export execDetailsEndCallback
func execDetailsEndCallback(id _Ctype_long, reqID _Ctype_int) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.ExecDetailsEnd(int(reqID))
	}
}

// *** EClientSocket ***

const timeoutMs = 2000

// IBClient represents an IB client socket
type IBClient struct {
	client *_Ctype_struct_IBClient
	id     _Ctype_long
}

// NewIBClient returns a new client socket with the given EWrapper callbacks
func NewIBClient(wrapper EWrapper) *IBClient {
	w.lock.Lock()
	next := w.next
	w.m[next] = wrapper
	w.next++
	w.lock.Unlock()

	return &IBClient{client: (_Cfunc_new_client)(next, timeoutMs), id: next}
}

// Connect attempts to connect to TWS/IBGateway on the given host/port and client ID
func (c *IBClient) Connect(host string, port, clientID int) bool {
	cHost := (_Cfunc_CString)(host)
	// This does NOT use std::string in underlying code, but should be not needed
	// when this call returns
	defer func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cHost))
	return bool(func(_cgo0 *_Ctype_struct_IBClient, _cgo1 *_Ctype_char, _cgo2 _Ctype_int, _cgo3 _Ctype_int) _Ctype__Bool {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_client_connect)(_cgo0, _cgo1, _cgo2, _cgo3);}(c.client, cHost, _Ctype_int(port), _Ctype_int(clientID)))
}

// Disconnect attempts to disconnect from TWS/IBGateway
func (c *IBClient) Disconnect() {
	func(_cgo0 *_Ctype_struct_IBClient) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_disconnect)(_cgo0);}(c.client)
}

// IsConnected returns the connection state of the client
func (c *IBClient) IsConnected() bool {
	return bool(func(_cgo0 *_Ctype_struct_IBClient) _Ctype__Bool {;	_cgoCheckPointer(_cgo0);	return (_Cfunc_client_is_connected)(_cgo0);}(c.client))
}

// ProcessMsg processes the next msg waiting on the client
func (c *IBClient) ProcessMsg() {
	func(_cgo0 *_Ctype_struct_IBClient) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_process_msg)(_cgo0);}(c.client)
}

// Delete frees the underlying CPP resources and removes the wrapper from the map
func (c *IBClient) Delete() {
	// First, get rid of the underlying socket to prevent callbacks
	func(_cgo0 *_Ctype_struct_IBClient) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_delete_client)(_cgo0);}(c.client)
	// Now remove the reference to the ewrapper from the map
	w.lock.Lock()
	delete(w.m, c.id)
	w.lock.Unlock()
}

// ReqAccountSummary requests the summaries for all accounts
func (c *IBClient) ReqAccountSummary(reqID int, group, tags string) {
	cGroup, cTags := (_Cfunc_CString)(group), (_Cfunc_CString)(tags)
	// NOTE: Since the underlying code uses C++ std::string, we know these will be copied by the
	// constructor and can therefore be safely deallocated when this function exits
	defer func() {
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cGroup))
		func(_cgo0 _cgo_unsafe.Pointer) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_free)(_cgo0);}(unsafe.Pointer(cTags))
	}()
	func(_cgo0 *_Ctype_struct_IBClient, _cgo1 _Ctype_int, _cgo2 *_Ctype_char, _cgo3 *_Ctype_char) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_req_account_summ)(_cgo0, _cgo1, _cgo2, _cgo3);}(c.client, _Ctype_int(reqID), cGroup, cTags)
}

// CancelAccountSummary cancels the account summary info
func (c *IBClient) CancelAccountSummary(reqID int) {
	func(_cgo0 *_Ctype_struct_IBClient, _cgo1 _Ctype_int) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_cancel_account_summ)(_cgo0, _cgo1);}(c.client, _Ctype_int(reqID))
}

// PlaceOrder places an order with the given contracts
func (c *IBClient) PlaceOrder(orderID OrderID, contract *Contract, order *Order) {
	func(_cgo0 *_Ctype_struct_IBClient, _cgo1 _Ctype_OrderId, _cgo2 *_Ctype_struct_Contract, _cgo3 *_Ctype_struct_Order) {;	_cgoCheckPointer(_cgo0);	_cgoCheckPointer(_cgo2);	_cgoCheckPointer(_cgo3);	(_Cfunc_client_place_order)(_cgo0, _cgo1, _cgo2, _cgo3);}(c.client, orderID, contract.contract, order.order)
}

// CancelOrder cancels the order with the given order id
func (c *IBClient) CancelOrder(orderID OrderID) {
	func(_cgo0 *_Ctype_struct_IBClient, _cgo1 _Ctype_OrderId) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_cancel_order)(_cgo0, _cgo1);}(c.client, orderID)
}

// ReqOpenOrders requests all open orders
func (c *IBClient) ReqOpenOrders() {
	func(_cgo0 *_Ctype_struct_IBClient) {;	_cgoCheckPointer(_cgo0);	(_Cfunc_client_req_open_orders)(_cgo0);}(c.client)
}

// ReqExecutions requests details on executed orders
func (c *IBClient) ReqExecutions(reqID int, filter *ExecutionFilter) {
	func(_cgo0 *_Ctype_struct_IBClient, _cgo1 _Ctype_int, _cgo2 *_Ctype_struct_ExecutionFilter) {;	_cgoCheckPointer(_cgo0);	_cgoCheckPointer(_cgo2);	(_Cfunc_client_req_executions)(_cgo0, _cgo1, _cgo2);}(c.client, _Ctype_int(reqID), filter.filter)
}
