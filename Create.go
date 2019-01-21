package main

import (
	"github.com/jereksel/go-sonatype-nexus/scripts"
)

func Create(conf Configuration, req CreateHostedMavenRepositoryRequest) error {

	uuid, err := uploadGroovyScript(conf, scripts.CreateHostedMavenScript)

	if err != nil {
		return err
	}

	_, err = invokeScript(conf, uuid, req)

	if err != nil {
		return err
	}

	return nil
}
