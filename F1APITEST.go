package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main3() {

	url := "https://api.formula1.com/6657193977244c13?d=account.formula1.com"
	method := "POST"

	payload := strings.NewReader(`{"solution":{"interrogation":{"st":162229509,"sr":1959639815,"cr":78830557},"version":"stable"},"error":null,"performance":{"interrogation":185}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
