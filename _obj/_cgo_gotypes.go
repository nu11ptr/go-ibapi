// Code generated by cmd/cgo; DO NOT EDIT.

package ibapi

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_Contract = _Ctype_struct_Contract

type _Ctype_IBClient = _Ctype_struct_IBClient

type _Ctype_OrderId = _Ctype_long

type _Ctype__Bool bool

type _Ctype_char int8

type _Ctype_int int32

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

type _Ctype_struct_Contract struct{}

type _Ctype_struct_IBClient struct{}

type _Ctype_ulong uint64

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_cmalloc(uint64(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ _cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ byte
var _cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ)

//go:cgo_unsafe_args
func _Cfunc_client_cancel_account_summ(p0 *_Ctype_struct_IBClient, p1 _Ctype_int) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_cancel_account_summ, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_connect
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_connect _cgo_b43a3ff11f6e_Cfunc_client_connect
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_connect byte
var _cgo_b43a3ff11f6e_Cfunc_client_connect = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_connect)

//go:cgo_unsafe_args
func _Cfunc_client_connect(p0 *_Ctype_struct_IBClient, p1 *_Ctype_char, p2 _Ctype_int, p3 _Ctype_int) (r1 _Ctype__Bool) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_connect, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_disconnect
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_disconnect _cgo_b43a3ff11f6e_Cfunc_client_disconnect
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_disconnect byte
var _cgo_b43a3ff11f6e_Cfunc_client_disconnect = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_disconnect)

//go:cgo_unsafe_args
func _Cfunc_client_disconnect(p0 *_Ctype_struct_IBClient) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_disconnect, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_is_connected
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_is_connected _cgo_b43a3ff11f6e_Cfunc_client_is_connected
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_is_connected byte
var _cgo_b43a3ff11f6e_Cfunc_client_is_connected = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_is_connected)

//go:cgo_unsafe_args
func _Cfunc_client_is_connected(p0 *_Ctype_struct_IBClient) (r1 _Ctype__Bool) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_is_connected, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_process_msg
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_process_msg _cgo_b43a3ff11f6e_Cfunc_client_process_msg
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_process_msg byte
var _cgo_b43a3ff11f6e_Cfunc_client_process_msg = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_process_msg)

//go:cgo_unsafe_args
func _Cfunc_client_process_msg(p0 *_Ctype_struct_IBClient) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_process_msg, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_client_req_account_summ
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_client_req_account_summ _cgo_b43a3ff11f6e_Cfunc_client_req_account_summ
var __cgofn__cgo_b43a3ff11f6e_Cfunc_client_req_account_summ byte
var _cgo_b43a3ff11f6e_Cfunc_client_req_account_summ = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_client_req_account_summ)

//go:cgo_unsafe_args
func _Cfunc_client_req_account_summ(p0 *_Ctype_struct_IBClient, p1 _Ctype_int, p2 *_Ctype_char, p3 *_Ctype_char) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_client_req_account_summ, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_contract_currency
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_currency _cgo_b43a3ff11f6e_Cfunc_contract_currency
var __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_currency byte
var _cgo_b43a3ff11f6e_Cfunc_contract_currency = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_contract_currency)

//go:cgo_unsafe_args
func _Cfunc_contract_currency(p0 *_Ctype_struct_Contract) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_contract_currency, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_contract_exchange
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_exchange _cgo_b43a3ff11f6e_Cfunc_contract_exchange
var __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_exchange byte
var _cgo_b43a3ff11f6e_Cfunc_contract_exchange = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_contract_exchange)

//go:cgo_unsafe_args
func _Cfunc_contract_exchange(p0 *_Ctype_struct_Contract) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_contract_exchange, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_contract_month
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_month _cgo_b43a3ff11f6e_Cfunc_contract_month
var __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_month byte
var _cgo_b43a3ff11f6e_Cfunc_contract_month = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_contract_month)

//go:cgo_unsafe_args
func _Cfunc_contract_month(p0 *_Ctype_struct_Contract) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_contract_month, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_contract_sec_type
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_sec_type _cgo_b43a3ff11f6e_Cfunc_contract_sec_type
var __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_sec_type byte
var _cgo_b43a3ff11f6e_Cfunc_contract_sec_type = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_contract_sec_type)

//go:cgo_unsafe_args
func _Cfunc_contract_sec_type(p0 *_Ctype_struct_Contract) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_contract_sec_type, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_contract_symbol
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_symbol _cgo_b43a3ff11f6e_Cfunc_contract_symbol
var __cgofn__cgo_b43a3ff11f6e_Cfunc_contract_symbol byte
var _cgo_b43a3ff11f6e_Cfunc_contract_symbol = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_contract_symbol)

//go:cgo_unsafe_args
func _Cfunc_contract_symbol(p0 *_Ctype_struct_Contract) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_contract_symbol, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_delete_client
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_delete_client _cgo_b43a3ff11f6e_Cfunc_delete_client
var __cgofn__cgo_b43a3ff11f6e_Cfunc_delete_client byte
var _cgo_b43a3ff11f6e_Cfunc_delete_client = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_delete_client)

//go:cgo_unsafe_args
func _Cfunc_delete_client(p0 *_Ctype_struct_IBClient) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_delete_client, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_delete_contract
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_delete_contract _cgo_b43a3ff11f6e_Cfunc_delete_contract
var __cgofn__cgo_b43a3ff11f6e_Cfunc_delete_contract byte
var _cgo_b43a3ff11f6e_Cfunc_delete_contract = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_delete_contract)

//go:cgo_unsafe_args
func _Cfunc_delete_contract(p0 *_Ctype_struct_Contract) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_delete_contract, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_free
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_free _cgo_b43a3ff11f6e_Cfunc_free
var __cgofn__cgo_b43a3ff11f6e_Cfunc_free byte
var _cgo_b43a3ff11f6e_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_new_client
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_new_client _cgo_b43a3ff11f6e_Cfunc_new_client
var __cgofn__cgo_b43a3ff11f6e_Cfunc_new_client byte
var _cgo_b43a3ff11f6e_Cfunc_new_client = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_new_client)

//go:cgo_unsafe_args
func _Cfunc_new_client(p0 _Ctype_long, p1 _Ctype_ulong) (r1 *_Ctype_struct_IBClient) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_new_client, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc_new_contract
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc_new_contract _cgo_b43a3ff11f6e_Cfunc_new_contract
var __cgofn__cgo_b43a3ff11f6e_Cfunc_new_contract byte
var _cgo_b43a3ff11f6e_Cfunc_new_contract = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc_new_contract)

//go:cgo_unsafe_args
func _Cfunc_new_contract(p0 *_Ctype_char, p1 *_Ctype_char, p2 *_Ctype_char, p3 *_Ctype_char, p4 *_Ctype_char) (r1 *_Ctype_struct_Contract) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc_new_contract, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
		_Cgo_use(p4)
	}
	return
}
//go:cgo_export_dynamic updateAccountTimeCallback
//go:linkname _cgoexp_b43a3ff11f6e_updateAccountTimeCallback _cgoexp_b43a3ff11f6e_updateAccountTimeCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_updateAccountTimeCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_updateAccountTimeCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_updateAccountTimeCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_updateAccountTimeCallback(p0 _Ctype_long, p1 *_Ctype_char) {
	updateAccountTimeCallback(p0, p1)
}
//go:cgo_export_dynamic errorCallback
//go:linkname _cgoexp_b43a3ff11f6e_errorCallback _cgoexp_b43a3ff11f6e_errorCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_errorCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_errorCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_errorCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_errorCallback(p0 _Ctype_long, p1 _Ctype_int, p2 _Ctype_int, p3 *_Ctype_char) {
	errorCallback(p0, p1, p2, p3)
}
//go:cgo_export_dynamic connectionClosedCallback
//go:linkname _cgoexp_b43a3ff11f6e_connectionClosedCallback _cgoexp_b43a3ff11f6e_connectionClosedCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_connectionClosedCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_connectionClosedCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_connectionClosedCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_connectionClosedCallback(p0 _Ctype_long) {
	connectionClosedCallback(p0)
}
//go:cgo_export_dynamic nextValidIDCallback
//go:linkname _cgoexp_b43a3ff11f6e_nextValidIDCallback _cgoexp_b43a3ff11f6e_nextValidIDCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_nextValidIDCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_nextValidIDCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_nextValidIDCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_nextValidIDCallback(p0 _Ctype_long, p1 _Ctype_OrderId) {
	nextValidIDCallback(p0, p1)
}
//go:cgo_export_dynamic accountSummaryCallback
//go:linkname _cgoexp_b43a3ff11f6e_accountSummaryCallback _cgoexp_b43a3ff11f6e_accountSummaryCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_accountSummaryCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_accountSummaryCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_accountSummaryCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_accountSummaryCallback(p0 _Ctype_long, p1 _Ctype_int, p2 *_Ctype_char, p3 *_Ctype_char, p4 *_Ctype_char, p5 *_Ctype_char) {
	accountSummaryCallback(p0, p1, p2, p3, p4, p5)
}
//go:cgo_export_dynamic accountSummaryEndCallback
//go:linkname _cgoexp_b43a3ff11f6e_accountSummaryEndCallback _cgoexp_b43a3ff11f6e_accountSummaryEndCallback
//go:cgo_export_static _cgoexp_b43a3ff11f6e_accountSummaryEndCallback
//go:nosplit
//go:norace
func _cgoexp_b43a3ff11f6e_accountSummaryEndCallback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_b43a3ff11f6e_accountSummaryEndCallback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_b43a3ff11f6e_accountSummaryEndCallback(p0 _Ctype_long, p1 _Ctype_int) {
	accountSummaryEndCallback(p0, p1)
}

//go:cgo_import_static _cgo_b43a3ff11f6e_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_b43a3ff11f6e_Cfunc__Cmalloc _cgo_b43a3ff11f6e_Cfunc__Cmalloc
var __cgofn__cgo_b43a3ff11f6e_Cfunc__Cmalloc byte
var _cgo_b43a3ff11f6e_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_b43a3ff11f6e_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_b43a3ff11f6e_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
