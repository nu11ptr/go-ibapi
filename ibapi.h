#ifndef IBAPI_H
#define IBAPI_H

#ifdef __cplusplus
#include "Contract.h"
#include "StdAfx.h"
#include "Order.h"
#include "EReaderOSSignal.h"
#include "EClientSocket.h"
#include "DefaultEWrapper.h"
#include "EReader.h"
#include "Execution.h"

#include "_obj/_cgo_export.h"
#include "stdio.h"

class IBClient : public DefaultEWrapper
{
    long wrapper_id;
    EReaderOSSignal signal;
    EClientSocket sock;
    EReader *reader;

  public:
    IBClient(long wrapper_id, unsigned long timeout);

    virtual ~IBClient();

    bool connect(const char *host, int port, int clientId);

    void disconnect()
    {
        sock.eDisconnect();
    }

    bool isConnected()
    {
        return sock.isConnected();
    }

    void processMsg();

    void reqAccountSummary(int reqId, const std::string groupName, const std::string tags);

    void cancelAccountSummary(int reqId);

    void placeOrder(OrderId orderId, Contract &contract, Order &order);

    void cancelOrder(OrderId orderId);

    void reqOpenOrders()
    {
        sock.reqOpenOrders();
    }

    void reqExecutions(int reqId, ExecutionFilter &filter)
    {
        sock.reqExecutions(reqId, filter);
    }

    // *** EWrapper ****

    void nextValidId(OrderId orderId)
    {
        nextValidIDCallback(wrapper_id, orderId);
    }

    void updateAccountTime(const std::string &timeStamp)
    {
        updateAccountTimeCallback(wrapper_id, (char *)timeStamp.c_str());
    }

    void error(int id, int errorCode, const std::string &errorString)
    {
        errorCallback(wrapper_id, id, errorCode, (char *)errorString.c_str());
    }

    void connectAck()
    {
        sock.startApi();
    }

    void connectionClosed()
    {
        connectionClosedCallback(wrapper_id);
    }

    void accountSummary(int reqId, const std::string &account, const std::string &tag,
                        const std::string &value, const std::string &currency)
    {
        accountSummaryCallback(wrapper_id, reqId, (char *)account.c_str(), (char *)tag.c_str(),
                               (char *)value.c_str(), (char *)currency.c_str());
    }

    void accountSummaryEnd(int reqId)
    {
        accountSummaryEndCallback(wrapper_id, reqId);
    }

    // ### Called by 'reqOpenOrders' ###

    void openOrder(OrderId orderId, const Contract &contract, const Order &order,
                   const OrderState &state)
    {
        openOrderCallback(wrapper_id, orderId, (Contract *)&contract, (Order *)&order);
    }

    void orderStatus(OrderId orderId, const std::string &status, double filled, double remaining,
                     double avgFillPrice, int permId, int parentId, double lastFillPrice,
                     int clientId, const std::string &whyHeld, double mktCapPrice)
    {
        orderStatusCallback(wrapper_id, orderId, (char *)status.c_str(), filled, remaining, avgFillPrice,
                            permId, parentId, lastFillPrice, clientId, (char *)whyHeld.c_str(), mktCapPrice);
    }

    void orderBound(long long orderId, int apiClientId, int apiOrderId)
    {
        orderBoundCallback(wrapper_id, orderId, apiClientId, apiOrderId);
    }

    void openOrderEnd()
    {
        openOrderEndCallback(wrapper_id);
    }

    // ### Called by 'reqExecutions' ###

    void execDetails(int reqId, const Contract &contract, const Execution &execution)
    {
        execDetailsCallback(wrapper_id, reqId, (Contract *)&contract, (Execution *)&execution);
    }

    void commissionReport(const CommissionReport &commissionReport) {}

    void execDetailsEnd(int reqId)
    {
        execDetailsEndCallback(wrapper_id, reqId);
    }
};

extern "C"
{
#else
typedef struct Contract Contract;
typedef struct Order Order;
typedef struct ExecutionFilter ExecutionFilter;
typedef struct Execution Execution;
typedef long OrderId;

typedef struct IBClient IBClient;

#include <stdbool.h>

#endif
    // *** Contract ***

    Contract *new_contract(const char *sym, const char *sec_type, const char *exch,
                           const char *contract_month, const char *curr);

    void delete_contract(Contract *contract);

    const char *contract_symbol(Contract *contract);

    const char *contract_sec_type(Contract *contract);

    const char *contract_exchange(Contract *contract);

    const char *contract_month(Contract *contract);

    const char *contract_currency(Contract *contract);

    // *** Order ***

    Order *new_order(int order_id, const char *action, const char *type, double qty, double price, const char *tif);

    void delete_order(Order *order);

    int order_id(Order *order);

    const char *order_action(Order *order);

    const char *order_type(Order *order);

    double order_qty(Order *order);

    double order_price(Order *order);

    const char *order_tif(Order *order);

    // *** ExecutionFilter ***

    ExecutionFilter *new_exec_filter(long clientId, const char *acctCode, const char *execTime,
                                     const char *symbol, const char *secType, const char *exchange,
                                     const char *side);

    void delete_exec_filter(ExecutionFilter *filter);

    // *** Execution ***

    const char *exec_id(Execution *exec);

    const char *exec_time(Execution *exec);

    const char *exec_account_num(Execution *exec);

    const char *exec_exchange(Execution *exec);

    const char *exec_side(Execution *exec);

    double exec_shares(Execution *exec);

    double exec_price(Execution *exec);

    double exec_avg_price(Execution *exec);

    long exec_order_id(Execution *exec);

    // *** EClientSocket ***

    IBClient *new_client(long wrapper_id, unsigned long timeout);

    void delete_client(IBClient *client);

    bool client_connect(IBClient *client, const char *host, int port, int clientId);

    void client_disconnect(IBClient *client);

    bool client_is_connected(IBClient *client);

    void client_process_msg(IBClient *client);

    void client_req_account_summ(IBClient *client, int req_id, const char *group, const char *tags);

    void client_cancel_account_summ(IBClient *client, int req_id);

    void client_place_order(IBClient *client, OrderId orderId, Contract *contract, Order *order);

    void client_cancel_order(IBClient *client, OrderId orderId);

    void client_req_open_orders(IBClient *client);

    void client_req_executions(IBClient *client, int reqId, ExecutionFilter *filter);

#ifdef __cplusplus
}
#endif

#endif
