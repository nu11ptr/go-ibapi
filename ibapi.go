package ibapi

//go:generate go tool cgo ibapi.go

// NOTE: Before building, ensure you've exported these variables:
// CGO_LDFLAGS=<path_to_ib_cpp_api>/libTwsSocketClient.so
// CGO_CPP_FLAGS=-I<path_to_ib_cpp_api>

/*
#cgo CXXFLAGS: -std=c++11

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

//export nextValidIDCallback
func nextValidIDCallback(id C.long, orderID C.OrderId) {
	if wrapper := findEWrapper(id); wrapper != nil {
		wrapper.NextValidId(orderID)
	}
}

// *** EClientSocket ***

// EClientSocket represents an IB client socket
type EClientSocket struct {
	sock *C.ClientSock
	id   C.long
}

// NewEClientSocket returns a new client socket with the given EWrapper callbacks
func NewEClientSocket(wrapper EWrapper) *EClientSocket {
	w.lock.Lock()
	next := w.next
	w.m[next] = wrapper
	w.next++
	w.lock.Unlock()

	sock := &EClientSocket{sock: C.new_client_sock(next), id: next}
	return sock
}

// EConnect attempts to connect to TWS/IBGateway on the given host/port and client ID
func (s *EClientSocket) EConnect(host string, port, clientID int) {
	cHost := C.CString(host)
	defer C.free(unsafe.Pointer(cHost))
	C.sock_econnect(s.sock, cHost, C.int(port), C.int(clientID))
}

// EDisconnect attempts to disconnect from TWS/IBGateway
func (s *EClientSocket) EDisconnect() {
	C.sock_edisconnect(s.sock)
}

// Delete frees the underlying CPP resources and removes the wrapper from the map
func (s *EClientSocket) Delete() {
	// First, get rid of the underlying socket to prevent callbacks
	C.delete_client_sock(s.sock)
	// Now remove the reference to the ewrapper from the map
	w.lock.Lock()
	delete(w.m, s.id)
	w.lock.Unlock()
}
