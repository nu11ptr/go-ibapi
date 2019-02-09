#ifndef IBAPI_H
#define IBAPI_H

#ifdef __cplusplus
#include "Contract.h"

extern "C"
{
#else
typedef struct Contract Contract;

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

#ifdef __cplusplus
}
#endif

#endif
