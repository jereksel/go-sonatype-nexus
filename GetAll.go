package main

import (
	"encoding/json"

	"github.com/jereksel/go-sonatype-nexus/scripts"
)

type getAllResult struct {
	Name string            `json:"name"`
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

func GetAll(conf Configuration) ([]Repository, error) {

	uuid, err := uploadGroovyScript(conf, scripts.GetAllScript)

	if err != nil {
		return nil, err
	}

	ret, err := invokeScript(conf, uuid, "")

	if err != nil {
		return nil, err
	}

	var results []getAllResult

	if err := json.Unmarshal([]byte(ret), &results); err != nil {
		return nil, err
	}

	repositories := make([]Repository, 0)

	for _, result := range results {

		if result.Type == "maven_hosted" {
			repositories = append(repositories, MavenHostedRepository{result.Name})
		} else if result.Type == "maven_proxy" {
			repositories = append(repositories, MavenProxyRepository{result.Name, result.Data["remoteUrl"]})
		} else {

		}

	}

	return repositories, nil

}
