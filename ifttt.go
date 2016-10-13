package ifttt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Values struct {
	Value1 interface{} `json:"value1"`
	Value2 interface{} `json:"value2"`
	Value3 interface{} `json:"value3"`
}

var (
	urlFormat = "https://maker.ifttt.com/trigger/%s/with/key/%s"
	makerKey  = os.Getenv("IFTTT_MAKER_KEY")
)

func SetKey(key string) {
	makerKey = key
}

func Deliver(action string, values Values) error {
	url := fmt.Sprintf(urlFormat, action, makerKey)
	body, err := json.Marshal(values)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
