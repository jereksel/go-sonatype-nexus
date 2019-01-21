package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type newScriptRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func uploadGroovyScript(conf Configuration, script string) (string, error) {

	uuid := uuid.New().String()

	scriptRequest := newScriptRequest{uuid, script, "groovy"}

	json, err := json.Marshal(scriptRequest)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", conf.url+"/service/rest/v1/script", bytes.NewReader(json))

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	req.SetBasicAuth(conf.login, conf.password)

	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)

	if err != nil {
		return "", err
	}

	return uuid, nil

}

func invokeScript(conf Configuration, scriptName string, body interface{}) (string, error) {

	bodyString, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	bodyReader := bytes.NewReader(bodyString)

	req, err := http.NewRequest("POST", conf.url+"/service/rest/v1/script/"+scriptName+"/run", bodyReader)

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	req.SetBasicAuth(conf.login, conf.password)

	req.Header.Set("Content-Type", "text/plain")

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	arr, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var dat map[string]string

	if err := json.Unmarshal(arr, &dat); err != nil {
		return "", err
	}

	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return "", errors.New(dat["result"])
	}

	return dat["result"], nil
}
