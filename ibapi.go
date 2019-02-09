package ibapi

// NOTE: Before building, ensure you've exported these variables:
// CGO_LDFLAGS=<path_to_ib_cpp_api>/libTwsSocketClient.so
// CGO_CPP_FLAGS=-I<path_to_ib_cpp_api>

/*
#include <stdlib.h>
#include "ibapi.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// *** Contract ***

// Contract represents a contract
type Contract struct {
	contract *C.struct_Contract
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
