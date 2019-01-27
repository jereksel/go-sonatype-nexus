package main

import (
	"github.com/jereksel/go-sonatype-nexus/scripts"
)

func CreateHostedMaven(conf Configuration, req CreateHostedMavenRepositoryRequest) error {
	return create(conf, scripts.CreateHostedMavenScript, req)
}

func CreateProxyMaven(conf Configuration, req CreateProxyMavenRepositoryRequest) error {
	return create(conf, scripts.CreateProxyMavenScript, req)
}

func CreateGroupMaven(conf Configuration, req CreateGroupMavenRepositoryRequest) error {
	return create(conf, scripts.CreateGroupMavenScript, req)
}

func create(conf Configuration, script string, req interface{}) error {

	uuid, err := uploadGroovyScript(conf, script)

	if err != nil {
		return err
	}

	_, err = invokeScript(conf, uuid, req)

	if err != nil {
		return err
	}

	return nil

}
