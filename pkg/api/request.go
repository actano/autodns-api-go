package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func MakeRequest(data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", Endpoint, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

    return body, nil
}
