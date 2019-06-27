package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	PINDO_TOKEN    = "eyJhbGciOiJub25lIn0.eyJpZCI6MjcsInJldm9rZWRfdG9rZW5fY291bnQiOjB9."
	PINDO_USERNAME = "opiumated"
	PINDO_EMAIL    = "all4usoro@gmail.com"
)

type MessageBody struct {
	To     string `json:"to"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

func SendSMS(to, text, sender string) {
	pindo_url := "http://api.pindo.io/v1/sms/"

	var message = MessageBody{
		To:     fmt.Sprintf("+25%s", to),
		Text:   text,
		Sender: "Ibisubizo",
	}

	buf, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	// payload := strings.NewReader(fmt.Sprintf("{"to" : "+250781234567", "text" : "Hello from Pindo","sender" : "Ibisubizo"}"))
	payload := bytes.NewBufferString(string(buf))
	log.Println("Payload")
	log.Println(payload)

	req, _ := http.NewRequest("POST", pindo_url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", PINDO_TOKEN))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res)
	fmt.Println(string(body))
}
