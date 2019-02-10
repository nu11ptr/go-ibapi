#include "ibapi.h"
#include "_obj/_cgo_export.h"

#include "EClientSocket.h"

// *** Contract ***

Contract *new_contract(const char *sym, const char *sec_type, const char *exch,
                       const char *contract_month, const char *curr)
{
    Contract *contract = new Contract();
    contract->symbol = std::string(sym);
    contract->secType = std::string(sec_type);
    contract->exchange = std::string(exch);
    contract->lastTradeDateOrContractMonth = std::string(contract_month);
    contract->currency = std::string(curr);
    return contract;
}

void delete_contract(Contract *contract)
{
    delete contract;
}

const char *contract_symbol(Contract *contract)
{
    return contract->symbol.c_str();
}

const char *contract_sec_type(Contract *contract)
{
    return contract->secType.c_str();
}

const char *contract_exchange(Contract *contract)
{
    return contract->exchange.c_str();
}

const char *contract_month(Contract *contract)
{
    return contract->lastTradeDateOrContractMonth.c_str();
}

const char *contract_currency(Contract *contract)
{
    return contract->currency.c_str();
}

// *** EClientSocket ***

ClientSock *new_client_sock(long wrapper_id)
{
    return new ClientSock(wrapper_id);
}

void delete_client_sock(ClientSock *sock)
{
    delete sock;
}

void sock_econnect(ClientSock *sock, const char *host, int port, int clientId)
{
    sock->sock.eConnect(host, port, clientId);
}

void sock_edisconnect(ClientSock *sock)
{
    sock->sock.eDisconnect();
}