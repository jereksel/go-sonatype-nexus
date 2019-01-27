package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	conf := getConf()

	removeAll(conf)

	hostedMavenName := "a-hosted-maven"

	if err := CreateHostedMaven(conf, CreateHostedMavenRepositoryRequest{hostedMavenName}); err != nil {
		panic(err)
	}

	proxyMavenName := "b-proxy-maven"

	if err := CreateProxyMaven(conf, CreateProxyMavenRepositoryRequest{proxyMavenName, "http://google.com"}); err != nil {
		panic(err)
	}

	body, err := GetAll(conf)

	if err != nil {
		panic(err)
	}

	exp := []Repository{
		MavenHostedRepository{hostedMavenName},
		MavenProxyRepository{proxyMavenName, "http://google.com"},
	}

	assert.Equal(t, exp, body)

}
