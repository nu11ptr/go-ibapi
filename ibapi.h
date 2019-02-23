#ifndef IBAPI_H
#define IBAPI_H

#ifdef __cplusplus
#include "Contract.h"
#include "StdAfx.h"
#include "EReaderOSSignal.h"
#include "EClientSocket.h"
#include "DefaultEWrapper.h"
#include "EReader.h"

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

    // *** EWrapper ****

    void nextValidId(OrderId orderId)
    {
        nextValidIDCallback(wrapper_id, orderId);
    }

    void updateAccountTime(const std::string &timeStamp);

    void error(int id, int errorCode, const std::string &errorString);

    void connectAck()
    {
        sock.startApi();
    }

    void connectionClosed();
};

extern "C"
{
#else
typedef struct Contract Contract;
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

    // *** EClientSocket ***

    IBClient *new_client(long wrapper_id, unsigned long timeout);

    void delete_client(IBClient *client);

    bool client_connect(IBClient *client, const char *host, int port, int clientId);

    void client_disconnect(IBClient *client);

    bool client_is_connected(IBClient *client);

    void client_process_msg(IBClient *client);

#ifdef __cplusplus
}
#endif

#endif
