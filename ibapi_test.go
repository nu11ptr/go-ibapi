package ibapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ibapi "gitlab.com/apisw/go-ibapi"
)

func TestContract(t *testing.T) {
	c := ibapi.NewContract("ES", "FUT", "GLOBEX", "202003", "USD")
	assert.Equal(t, "ES", c.Symbol())
	assert.Equal(t, "FUT", c.SecType())
	assert.Equal(t, "GLOBEX", c.Exchange())
	assert.Equal(t, "202003", c.ContractMonth())
	assert.Equal(t, "USD", c.Currency())
}
