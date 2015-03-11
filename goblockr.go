package goblockr

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var CoinList = []string{"btc", "tbtc", "ltc", "dgc", "grk", "ppc", "mec"}

type API struct {
	coin string
	httpTransport *http.Transport
}

func NewAPI(coin string) *API {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &API{coin, tr}
}

func (this *API) Call(action string, params map[string]string) ([]byte, error) {

	var err error
	var res *http.Response

	values := ""
	for key, val := range params {
		values += "&" + key + "=" + val
	}

	res, err = http.Get("http://" + this.coin + ".blockr.io/api/v1" + action + "?" + values)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, err
}

func (this *API) SetCoin(coin string) bool{
	coin = strings.ToLower(coin)
	for _, c := range CoinList {
		if c == coin {
			this.coin = coin
			return true
		}
	}
	return false
}

func (this *API) CoinInfo() (interface{}, error) {
	dataStream, err := this.Call("/coin/info", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) ExchangeRate() (interface{}, error) {
dataStream, err := this.Call("/exchangerate/current", nil)
data := map[string]interface{}{}
json.Unmarshal(dataStream, &data)
return data, err
}

func (this *API) BlockInfo(blockHash string) (interface{}, error) {
	dataStream, err := this.Call("/block/info/" + blockHash, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) BlockTxs(blockHash string) (interface{}, error) {
	dataStream, err := this.Call("/block/txs/" + blockHash, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) BlockRaw(blockHash string) (interface{}, error) {
	dataStream, err := this.Call("/block/raw/" + blockHash, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) TxInfo(hash, format string) (interface{}, error) {
	var params = map[string]string{
		"amount_format": format,
	}
	dataStream, err := this.Call("/tx/info/" + hash, params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) TxRaw(hash string) (interface{}, error) {
	dataStream, err := this.Call("/tx/raw/" + hash, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) UnconfirmedTxInfo(hash string) (interface{}, error) {
	dataStream, err := this.Call("/zerotx/info/" + hash, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) AddressInfo(address string, confirmations int, format string) (interface{}, error) {
	var params = map[string]string{
		"confirmations": strconv.Itoa(confirmations),
		"amount_format": format,
	}
	dataStream, err := this.Call("/address/info/" + address, params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) AddressBalance(address string, confirmations int) (interface{}, error) {
	var params = map[string]string{
		"confirmations": strconv.Itoa(confirmations),
	}
	dataStream, err := this.Call("/address/balance/" + address, params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}


func (this *API) AddressTxs(address string) (interface{}, error) {
	dataStream, err := this.Call("/address/txs/" + address, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) AddressUnspent(address string, unconfirmed, multisigs bool) (interface{}, error) {
	unconfirmed_str := "0"
	if unconfirmed {
		unconfirmed_str = "1"
	}
	multisigs_str := "0"
	if multisigs {
		multisigs_str = "1"
	}
	var params = map[string]string{
		"unconfirmed": unconfirmed_str,
		"multisigs": multisigs_str,
	}
	dataStream, err := this.Call("/address/unspent/" + address, params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) AddressUnconfirmed(address string) (interface{}, error) {
	dataStream, err := this.Call("/address/unconfirmed/" + address, nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

