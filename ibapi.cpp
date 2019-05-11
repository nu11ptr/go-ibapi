#include "ibapi.h"
#include <cstring>

// *** Contract ***

Contract *new_contract(const char *sym, const char *sec_type, const char *exch,
                       const char *contract_month, const char *curr)
{
    Contract *contract = new Contract();
    contract->symbol = sym;
    contract->secType = sec_type;
    contract->exchange = exch;
    contract->lastTradeDateOrContractMonth = contract_month;
    contract->currency = curr;
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

// *** Order ***

Order *new_order(int order_id, const char *action, const char *type, double qty, double price, const char *tif)
{
    Order *order = new Order();
    order->orderId = order_id;
    order->action = action;
    order->orderType = type;
    order->totalQuantity = qty;
    if (strcmp(type, "LMT") == 0 || strcmp(type, "lmt") == 0)
    {
        order->lmtPrice = price;
    }
    else if (strcmp(type, "STP") == 0 || strcmp(type, "stp") == 0)
    {
        order->auxPrice = price;
    }
    order->tif = tif;
    order->transmit = true;
    return order;
}

void delete_order(Order *order)
{
    delete order;
}

int order_id(Order *order)
{
    return order->orderId;
}

const char *order_action(Order *order)
{
    return order->action.c_str();
}

const char *order_type(Order *order)
{
    return order->orderType.c_str();
}

double order_qty(Order *order)
{
    return order->totalQuantity;
}

double order_price(Order *order)
{
    const char *type = order->orderType.c_str();
    if (strcmp(type, "LMT") == 0 || strcmp(type, "lmt") == 0)
    {
        return order->lmtPrice;
    }
    else if (strcmp(type, "STP") == 0 || strcmp(type, "stp") == 0)
    {
        return order->auxPrice;
    }
    return 0.0;
}

const char *order_tif(Order *order)
{
    return order->tif.c_str();
}

// *** ExecutionFilter ***

ExecutionFilter *new_exec_filter(long clientId, const char *acctCode, const char *execTime,
                                 const char *symbol, const char *secType, const char *exchange,
                                 const char *side)
{
    ExecutionFilter *filter = new ExecutionFilter();
    filter->m_clientId = clientId;
    filter->m_acctCode = acctCode;
    filter->m_time = execTime;
    filter->m_symbol = symbol;
    filter->m_secType = secType;
    filter->m_exchange = exchange;
    filter->m_side = side;
    return filter;
}

void delete_exec_filter(ExecutionFilter *filter)
{
    delete filter;
}

// *** Execution ***

const char *exec_id(Execution *exec)
{
    return exec->execId.c_str();
}

const char *exec_time(Execution *exec)
{
    return exec->time.c_str();
}

const char *exec_account_num(Execution *exec)
{
    return exec->acctNumber.c_str();
}

const char *exec_exchange(Execution *exec)
{
    return exec->exchange.c_str();
}

const char *exec_side(Execution *exec)
{
    return exec->side.c_str();
}

double exec_shares(Execution *exec)
{
    return exec->shares;
}

double exec_price(Execution *exec)
{
    return exec->price;
}

double exec_avg_price(Execution *exec)
{
    return exec->avgPrice;
}

long exec_order_id(Execution *exec)
{
    return exec->orderId;
}

// *** IBClient ***

IBClient::IBClient(long wrapper_id, unsigned long timeout)
    : wrapper_id(wrapper_id), signal(timeout), sock(this, &signal), reader(0)
{
}

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

void client_req_account_summ(IBClient *client, int req_id, const char *group, const char *tags)
{
    client->reqAccountSummary(req_id, group, tags);
}

void client_cancel_account_summ(IBClient *client, int req_id)
{
    client->cancelAccountSummary(req_id);
}

void client_place_order(IBClient *client, OrderId orderId, Contract *contract, Order *order)
{
    client->placeOrder(orderId, *contract, *order);
}

void client_cancel_order(IBClient *client, OrderId orderId)
{
    client->cancelOrder(orderId);
}

void client_req_open_orders(IBClient *client)
{
    client->reqOpenOrders();
}

void client_req_executions(IBClient *client, int reqId, ExecutionFilter *filter)
{
    client->reqExecutions(reqId, *filter);
}
