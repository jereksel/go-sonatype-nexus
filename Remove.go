package main

import (
	"encoding/json"
	"fmt"

	"github.com/jereksel/go-sonatype-nexus/scripts"
)

type removeRequest struct {
	Name string `json:"name"`
}

type removeResult struct {
	Status bool `json:"status"`
}

func Remove(conf Configuration, repositoryName string) error {

	uuid, err := uploadGroovyScript(conf, scripts.RemoveScript)

	if err != nil {
		return err
	}

	request := removeRequest{repositoryName}

	ret, err := invokeScript(conf, uuid, request)

	if err != nil {
		return err
	}

	var result removeResult

	if err := json.Unmarshal([]byte(ret), &result); err != nil {
		return err
	}

	if !result.Status {
		return fmt.Errorf("Repository '%s' does not exist", repositoryName)
	}

	return nil
}
