package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.Handle("/", Router())
}

func TestConnection(t *testing.T) {
	var err error
	url := "https://ropsten.coinapp.io"
	conn, err = ethclient.Dial(url)
	assert.Nil(t, err)
}

func TestFormatDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("120984932420357242390102000000000000000", 10)
	tokenCorrected := BigIntDecimal(number, 18)
	assert.Equal(t, "120984932420357242390.102", tokenCorrected)
}

func TestBalanceCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, "10000.0", rr.Body.String(), "should be balance")
}

func TestTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d BalanceResponse
	json.Unmarshal(rr.Body.Bytes(), &d)

	assert.Equal(t, "DreamTeam Token", d.Name, "should be token name")
	assert.Equal(t, "0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", d.Wallet, "should be wallet address")
	assert.Equal(t, uint8(0x6), d.Decimals, "should be decimals")
	assert.Equal(t, "DTT", d.Symbol, "should be symbol")
	assert.Equal(t, "10000.0", d.Balance, "should be Token balance")
	assert.Equal(t, "50.0", d.EthBalance, "should be ETH balance")
}

func TestMainnetConnection(t *testing.T) {
	var err error
	url := "https://eth.coinapp.io"
	conn, err = ethclient.Dial(url)
	assert.Nil(t, err)
}

func TestMainnetTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xd26114cd6EE289AccF82350c8d8487fedB8A0C07/0x42d4722b804585cdf6406fa7739e794b0aa8b1ff", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d BalanceResponse
	json.Unmarshal(rr.Body.Bytes(), &d)

	assert.Equal(t, "OMGToken", d.Name, "should be token name")
	assert.Equal(t, "0x42d4722b804585cdf6406fa7739e794b0aa8b1ff", d.Wallet, "should be wallet address")
	assert.Equal(t, uint8(0x12), d.Decimals, "should be decimals")
	assert.Equal(t, "OMG", d.Symbol, "should be symbol")
	assert.Equal(t, "600000.0", d.Balance, "should be Token balance")
}

func TestMainnetEOSTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0/0xbfaa1a1ea534d35199e84859975648b59880f639", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d BalanceResponse
	json.Unmarshal(rr.Body.Bytes(), &d)

	assert.Equal(t, "", d.Name, "should be token name")
	assert.Equal(t, "0xbfaa1a1ea534d35199e84859975648b59880f639", d.Wallet, "should be wallet address")
	assert.Equal(t, uint8(0x12), d.Decimals, "should be decimals")
	assert.Equal(t, "EOS", d.Symbol, "should be symbol")
	assert.Equal(t, "9200000.0", d.Balance, "should be Token balance")
}
