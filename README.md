# go-ibapi

## Build TWS Shared Library (Linux)

NOTE: Clong IB TWS requires authorization of your github ID from IB

```bash
git clone git@github.com:InteractiveBrokers/tws-api.git
cd tws-api
git checkout 9.74.01
cd tws-api/source/cppclient/client
make -f Makefile.linux
```

## Build Go Library (Linux)

```bash
export CGO_LDFLAGS=<path_to_ib_cpp_api>/libTwsSocketClient.so
export CGO_CPPFLAGS=-I<path_to_ib_cpp_api>
go generate ./...
go build ./...
```
