package client

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct{}

type ClientInterface interface {
	NewRequest(method string, url string, body io.Reader) []byte
}

func NewClient() ClientInterface {
	return &Client{}
}

func (*Client) NewRequest(method string, url string, body io.Reader) []byte {

	client := &http.Client{}

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalln("Error from the get pets API - ", err.Error())
	}

	//Read the response body
	resp, err := client.Do(r)
	if err != nil {
		log.Fatalln("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return respBody
}
