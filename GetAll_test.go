package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizationsService_One(t *testing.T) {

	conf := getConf()

	removeAll(conf)
	hostedMavenName := uuid.New().String()

	if err := Create(conf, CreateHostedMavenRepositoryRequest{hostedMavenName}); err != nil {
		panic(err)
	}

	body, err := GetAll(conf)

	if err != nil {
		panic(err)
	}

	exp := []Repository{Repository{hostedMavenName, "maven2", "hosted"}}

	assert.Equal(t, exp, body, "The two words should be the same.")

}
