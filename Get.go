package main

import (
	"encoding/json"
	"fmt"

	"github.com/jereksel/go-sonatype-nexus/scripts"
)

type getResult struct {
	Name string            `json:"name"`
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

func (result getResult) convert() (Repository, error) {

	if result.Type == "maven_hosted" {
		return MavenHostedRepository{result.Name}, nil
	} else if result.Type == "maven_proxy" {
		return MavenProxyRepository{result.Name, result.Data["remoteUrl"]}, nil
	} else if result.Type == "maven_group" {
		membersString := result.Data["members"]
		var members []string

		if err := json.Unmarshal([]byte(membersString), &members); err != nil {
			return nil, err
		}

		return MavenGroupRepository{result.Name, members}, nil
	} else {
		return nil, fmt.Errorf("Unknown result type: %s", result.Type)
	}

}

func Get(conf Configuration, name string) (*Repository, error) {

	uuid, err := uploadGroovyScript(conf, scripts.GetAllScript)

	if err != nil {
		return nil, err
	}

	ret, err := invokeScript(conf, uuid, map[string]string{
		"id": name,
	})

	if err != nil {
		return nil, err
	}

	if len(ret) == 0 || ret == "null" {
		return nil, nil
	}

	var result getResult

	if err := json.Unmarshal([]byte(ret), &result); err != nil {
		return nil, err
	}

	repo, err := result.convert()

	if err != nil {
		return nil, err
	}

	if repo == nil {
		return nil, nil
	}

	return &repo, nil

}

func GetAll(conf Configuration) ([]Repository, error) {

	uuid, err := uploadGroovyScript(conf, scripts.GetAllScript)

	if err != nil {
		return nil, err
	}

	var empty struct{}

	ret, err := invokeScript(conf, uuid, empty)

	if err != nil {
		return nil, err
	}

	var results []getResult

	if err := json.Unmarshal([]byte(ret), &results); err != nil {
		return nil, err
	}

	repositories := make([]Repository, len(results))

	for i, result := range results {

		repo, err := result.convert()

		if err != nil {
			return nil, err
		}

		repositories[i] = repo

	}

	return repositories, nil

}
