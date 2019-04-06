# go-ibapi

## Build TWS Shared Library (Linux)

1. `git clone git@github.com:InteractiveBrokers/tws-api.git` # requires authorization of your github ID from IB
2. `git checkout 9.74.01`
3. `cd tws-api/source/cppclient/client`
4. `make -f Makefile.linux`

## Build Go Library (Linux)

1. `export CGO_LDFLAGS=<path_to_ib_cpp_api>/libTwsSocketClient.so`
2. `export CGO_CPPFLAGS=-I<path_to_ib_cpp_api>`
3. `go generate ./...`
4. `go build ./...`
