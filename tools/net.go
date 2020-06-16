package tools

import (
	"bytes"
	"errors"
	json "github.com/json-iterator/go"
	"log"
	"net/http"
)


func Get(url string, params map[string]string, headers map[string]string, target interface{}) error {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return json.NewDecoder(resp.Body).Decode(target)
	}
	return nil
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string, target interface{}) error {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return json.NewDecoder(resp.Body).Decode(target)
	}
	return nil
}
