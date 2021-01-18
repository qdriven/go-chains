package service

import "net/http"

func GetEthTestCoin(address string) (*http.Response, error) {
	url := "https://faucet.ropsten.be/donate/"
	resp, err := http.Get(url + address)
	return resp, err
}
