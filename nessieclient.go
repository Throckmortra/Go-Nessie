package gonessie

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var BaseUrl = "http://api.reimaginebanking.com/"

type NessieClient struct {
	client   *http.Client
	apiKey   string
	payLoad  io.Reader
	response interface{}
	raw      *http.Response
}

func NesClient() *NessieClient {
	_, err := url.Parse(BaseUrl)
	if err != nil {
		panic("error parsing baseUrl")
	}

	clientInstance := &NessieClient{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	clientInstance.Client = client
	return clientInstance
}

func (r *NessieClient) SetKey(key string) {
	r.apiKey = "key=" + key
}

func (r *NessieClient) Account() Accounts {
	//account := &Account{id: "ayy"}
	account := Accounts{}

	fmt.Println(BaseUrl + "accounts?" + r.ApiKey)

	resp, err := r.client.Get(BaseUrl + "accounts?" + r.apiKey)
	if err != nil {
		fmt.Println("error1: " + err.Error())
	} else {
		fmt.Println("Status: ayyy " + resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error2: " + err.Error())
	}
	fmt.Printf("%s\n", string(contents))

	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		fmt.Println("error3: " + err.Error())
	}

	fmt.Println("length of account array: " + string(len(account)))

	//err = json.Unmarshal(test, &account)
	fmt.Println("object id: " + account[1].GetNickname() + " " + account[1].Nickname)
	defer resp.Body.Close()
	return account
}
