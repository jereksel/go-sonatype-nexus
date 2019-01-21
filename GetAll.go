package main

import (
	"encoding/json"

	"github.com/jereksel/go-sonatype-nexus/scripts"
)

type Repository struct {
	Name   string
	Format string
	Type   string
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

	var repos []Repository

	if err := json.Unmarshal([]byte(ret), &repos); err != nil {
		return nil, err
	}

	return repos, nil

}
