package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {

	url := "http://localhost:9091/api/list-meetings/upcoming"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("X-ACCESS-TOKEN", "HUJ/NBgLahavyTTptOU8p7vF5JFf8u3WEibA31+bJQNF")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-API-KEY", "AAhzclm9jHdypqdmEQx")
	req.Header.Add("X-API-NAME", "api_thirdparty")
	req.Header.Add("X-ACCESS-TOKEN", "bIxp+xMgTnD4p4p5JrxyEaokx6ZK2PS2xovklsDtrty0")

	req.Header.Set("Content-Type", writer.FormDataContentType())
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
