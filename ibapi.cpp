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

// *** IBClient ***

IBClient::IBClient(long wrapper_id, unsigned long timeout)
    : wrapper_id(wrapper_id), signal(timeout), sock(this, &signal), reader(0) {}

IBClient::~IBClient()
{
    if (reader)
    {
        delete reader;
    }
}

bool IBClient::connect(const char *host, int port, int clientId)
{
    auto connected = sock.eConnect(host, port, clientId);
    if (connected)
    {
        reader = new EReader(&sock, &signal);
        reader->start();
    }
    return connected;
}

void IBClient::processMsg()
{
    if (reader)
    {
        signal.waitForSignal();
        reader->processMsgs();
    }
}

void IBClient::reqAccountSummary(int reqId)
{
    sock.reqAccountSummary(reqId, "All", "");
}

void IBClient::cancelAccountSummary(int reqId)
{
    sock.cancelAccountSummary(reqId);
}

// *** EWrapper ***

// *** C API ***

IBClient *new_client(long wrapper_id, unsigned long timeout)
{
    return new IBClient(wrapper_id, timeout);
}

void delete_client(IBClient *client)
{
    delete client;
}

bool client_connect(IBClient *client, const char *host, int port, int clientId)
{
    return client->connect(host, port, clientId);
}

void client_disconnect(IBClient *client)
{
    client->disconnect();
}

bool client_is_connected(IBClient *client)
{
    return client->isConnected();
}

void client_process_msg(IBClient *client)
{
    client->processMsg();
}

void client_req_account_summ(IBClient *client, int req_id)
{
    client->reqAccountSummary(req_id);
}

void client_cancel_account_summ(IBClient *client, int req_id)
{
    client->cancelAccountSummary(req_id);
}
