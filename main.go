package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	toxiproxy "github.com/Shopify/toxiproxy/client"
	"net/http"
)

type ViaCepResponse struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
}

func (*ViaCepResponse) String() {

}

func SearchZipCode(host string, code string) (*ViaCepResponse, error) {
	client := http.DefaultClient
	response, err := client.Get(fmt.Sprintf("http://%s/ws/%s/json", host, code))
	if err != nil {
		return nil, err
	}
	result := new(ViaCepResponse)
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(response.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf.Bytes(), result); err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	tx_client := toxiproxy.NewClient("localhost:8474")
	proxy, err := tx_client.CreateProxy("viacep", "0.0.0.0:1234", "viacep.com.br:80")
	defer proxy.Delete()
	proxy.AddToxic("", "latency", "", 1, toxiproxy.Attributes{
		"latency": 1000,
	})
	if err != nil {
		panic(err)
	}
	response, err := SearchZipCode(proxy.Listen, "01001000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}
