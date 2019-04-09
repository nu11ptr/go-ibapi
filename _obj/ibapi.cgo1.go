// Code generated by cmd/cgo; DO NOT EDIT.

//line ibapi.go:1:1
package ibapi

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
	contract * /*line :25:12*/_Ctype_struct_Contract /*line :25:22*/
}

// NewContract creates a new contract
func NewContract(sym, secType, exch, contractMonth, currency string) *Contract {
	cSym, cSecType, cExch, cContractMonth, cCurrency := ( /*line :30:54*/_Cfunc_CString /*line :30:62*/)(sym), ( /*line :30:70*/_Cfunc_CString /*line :30:78*/)(secType),
		( /*line :31:3*/_Cfunc_CString /*line :31:11*/)(exch), ( /*line :31:20*/_Cfunc_CString /*line :31:28*/)(contractMonth), ( /*line :31:46*/_Cfunc_CString /*line :31:54*/)(currency)
	defer func() {
		func() { _cgo0 := /*line :33:10*/unsafe.Pointer(cSym); _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }()
		func() { _cgo0 := /*line :34:10*/unsafe.Pointer(cSecType); _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }()
		func() { _cgo0 := /*line :35:10*/unsafe.Pointer(cExch); _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }()
		func() { _cgo0 := /*line :36:10*/unsafe.Pointer(cContractMonth); _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }()
		func() { _cgo0 := /*line :37:10*/unsafe.Pointer(cCurrency); _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }()
	}()
	c := &Contract{contract: ( /*line :39:27*/_Cfunc_new_contract /*line :39:40*/)(cSym, cSecType, cExch, cContractMonth, cCurrency)}
	runtime.SetFinalizer(c, deleteContract)
	return c
}

func deleteContract(c *Contract) {
	func() { _cgo0 := /*line :45:20*/c.contract; _cgoCheckPointer(_cgo0); _Cfunc_delete_contract(_cgo0); }()
}

// Symbol returns the symbol from the contract
func (c *Contract) Symbol() string {
	return ( /*line :50:9*/_Cfunc_GoString /*line :50:18*/)(func() *_Ctype_char{ _cgo0 := /*line :50:38*/c.contract; _cgoCheckPointer(_cgo0); return _Cfunc_contract_symbol(_cgo0); }())
}

// SecType returns the security type of the contract
func (c *Contract) SecType() string {
	return ( /*line :55:9*/_Cfunc_GoString /*line :55:18*/)(func() *_Ctype_char{ _cgo0 := /*line :55:40*/c.contract; _cgoCheckPointer(_cgo0); return _Cfunc_contract_sec_type(_cgo0); }())
}

// Exchange returns the exchange of the contract
func (c *Contract) Exchange() string {
	return ( /*line :60:9*/_Cfunc_GoString /*line :60:18*/)(func() *_Ctype_char{ _cgo0 := /*line :60:40*/c.contract; _cgoCheckPointer(_cgo0); return _Cfunc_contract_exchange(_cgo0); }())
}

// ContractMonth returns the contract month of the futures contract
func (c *Contract) ContractMonth() string {
	return ( /*line :65:9*/_Cfunc_GoString /*line :65:18*/)(func() *_Ctype_char{ _cgo0 := /*line :65:37*/c.contract; _cgoCheckPointer(_cgo0); return _Cfunc_contract_month(_cgo0); }())
}

// Currency returns the currency used for the contract
func (c *Contract) Currency() string {
	return ( /*line :70:9*/_Cfunc_GoString /*line :70:18*/)(func() *_Ctype_char{ _cgo0 := /*line :70:40*/c.contract; _cgoCheckPointer(_cgo0); return _Cfunc_contract_currency(_cgo0); }())
}

// *** EWrapper ***

type wrappers struct {
	m    map[ /*line :76:11*/_Ctype_long /*line :76:17*/]EWrapper
	next  /*line :77:7*/_Ctype_long /*line :77:13*/
	lock sync.Mutex
}

var (
	w = wrappers{m: make(map[ /*line :82:27*/_Ctype_long /*line :82:33*/]EWrapper, 64)}
)

// OrderID represents an IB order ID
type OrderID =  /*line :86:16*/_Ctype_OrderId /*line :86:25*/

// EWrapper represesnts an interface of IB callbacks
type EWrapper interface {
	NextValidId(orderID OrderID)

	UpdateAccountTime(timeStamp string)

	Error(id, errorCode int, errorStr string)

	ConnectionClosed()

	AccountSummary(reqID int, account, tag, value string)

	AccountSummaryEnd(reqID int)
}

func findEWrapper(id  /*line :103:22*/_Ctype_long /*line :103:28*/) EWrapper {
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
func updateAccountTimeCallback(id  /*line :115:35*/_Ctype_long /*line :115:41*/, timeStamp * /*line :115:54*/_Ctype_char /*line :115:60*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.UpdateAccountTime(( /*line :117:29*/_Cfunc_GoString /*line :117:38*/)(timeStamp))
	}
}

//export errorCallback
func errorCallback(id  /*line :122:23*/_Ctype_long /*line :122:29*/, errID, errCode  /*line :122:46*/_Ctype_int /*line :122:51*/, errStr * /*line :122:61*/_Ctype_char /*line :122:67*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.Error(int(errID), int(errCode), ( /*line :124:43*/_Cfunc_GoString /*line :124:52*/)(errStr))
	}
}

//export connectionClosedCallback
func connectionClosedCallback(id  /*line :129:34*/_Ctype_long /*line :129:40*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.ConnectionClosed()
	}
}

//export nextValidIDCallback
func nextValidIDCallback(id  /*line :136:29*/_Ctype_long /*line :136:35*/, orderID  /*line :136:45*/_Ctype_OrderId /*line :136:54*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.NextValidId(orderID)
	}
}

//export accountSummaryCallback
func accountSummaryCallback(id  /*line :143:32*/_Ctype_long /*line :143:38*/, reqID  /*line :143:46*/_Ctype_int /*line :143:51*/, account, tag, value, currency * /*line :143:84*/_Ctype_char /*line :143:90*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummary(int(reqID), ( /*line :145:38*/_Cfunc_GoString /*line :145:47*/)(account), ( /*line :145:59*/_Cfunc_GoString /*line :145:68*/)(tag), ( /*line :145:76*/_Cfunc_GoString /*line :145:85*/)(currency))
	}
}

//export accountSummaryEndCallback
func accountSummaryEndCallback(id  /*line :150:35*/_Ctype_long /*line :150:41*/, reqID  /*line :150:49*/_Ctype_int /*line :150:54*/) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.AccountSummaryEnd(int(reqID))
	}
}

// *** EClientSocket ***

const timeoutMs = 2000

// IBClient represents an IB client socket
type IBClient struct {
	client * /*line :162:10*/_Ctype_struct_IBClient /*line :162:20*/
	id      /*line :163:9*/_Ctype_long /*line :163:15*/
}

// NewIBClient returns a new client socket with the given EWrapper callbacks
func NewIBClient(wrapper EWrapper) *IBClient {
	w.lock.Lock()
	next := w.next
	w.m[next] = wrapper
	w.next++
	w.lock.Unlock()

	return &IBClient{client: ( /*line :174:27*/_Cfunc_new_client /*line :174:38*/)(next, timeoutMs), id: next}
}

// Connect attempts to connect to TWS/IBGateway on the given host/port and client ID
func (c *IBClient) Connect(host string, port, clientID int) bool {
	cHost := ( /*line :179:11*/_Cfunc_CString /*line :179:19*/)(host)
	defer func() func() { _cgo0 := /*line :180:15*/unsafe.Pointer(cHost); return func() { _cgoCheckPointer(_cgo0); _Cfunc_free(_cgo0); }}()()
	return bool(func() _Ctype__Bool{ _cgo0 := /*line :181:31*/c.client; var _cgo1 *_Ctype_char = /*line :181:41*/cHost; var _cgo2 _Ctype_int = _Ctype_int(port); var _cgo3 _Ctype_int = _Ctype_int(clientID); _cgoCheckPointer(_cgo0); return _Cfunc_client_connect(_cgo0, _cgo1, _cgo2, _cgo3); }())
}

// Disconnect attempts to disconnect from TWS/IBGateway
func (c *IBClient) Disconnect() {
	func() { _cgo0 := /*line :186:22*/c.client; _cgoCheckPointer(_cgo0); _Cfunc_client_disconnect(_cgo0); }()
}

// IsConnected returns the connection state of the client
func (c *IBClient) IsConnected() bool {
	return bool(func() _Ctype__Bool{ _cgo0 := /*line :191:36*/c.client; _cgoCheckPointer(_cgo0); return _Cfunc_client_is_connected(_cgo0); }())
}

// ProcessMsg processes the next msg waiting on the client
func (c *IBClient) ProcessMsg() {
	func() { _cgo0 := /*line :196:23*/c.client; _cgoCheckPointer(_cgo0); _Cfunc_client_process_msg(_cgo0); }()
}

// Delete frees the underlying CPP resources and removes the wrapper from the map
func (c *IBClient) Delete() {
	// First, get rid of the underlying socket to prevent callbacks
	func() { _cgo0 := /*line :202:18*/c.client; _cgoCheckPointer(_cgo0); _Cfunc_delete_client(_cgo0); }()
	// Now remove the reference to the ewrapper from the map
	w.lock.Lock()
	delete(w.m, c.id)
	w.lock.Unlock()
}

func (c *IBClient) ReqAccountSummary() {

}
