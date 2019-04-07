# go-ibapi

Go bindings for the IB TWS C++ API Client. Note that these are in no way complete, but simply an abstracted wrapper over a portion of the API for our specific needs.

## Build TWS Shared Library (Linux)

This step is required before the library can be used.

NOTE: Cloning IB TWS requires authorization of your github ID from IB

```bash
git clone git@github.com:InteractiveBrokers/tws-api.git
cd tws-api
git checkout 9.74.01
cd tws-api/source/cppclient/client
make -f Makefile.linux
ln -s $PWD <path_to_your_project>/tws-api
```

# Install

NOTE: Ensure `Build TWS Shared Library` step is done before trying this

NOTE 2: The `_obj` folder contains generated CGo code for Go 1.12.1. If you are using a different version, you may need to clone and regenerate before using.

```bash
export CGO_LDFLAGS=tws-api/libTwsSocketClient.so
export CGO_CPPFLAGS=-Itws-api
go get -u github.com/nu11ptr/go-ibapi
```

## Regenerate CGO Support Files (_obj folder)

This is only required if you are using a Go version other than 1.12.x

```bash
export CGO_LDFLAGS=tws-api/libTwsSocketClient.so
export CGO_CPPFLAGS=-Itws-api
go generate ./...
```
