package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	conf := getConf()

	removeAll(conf)

	repositoryName := uuid.New().String()

	createRequest := CreateHostedMavenRepositoryRequest{repositoryName}

	err := CreateHostedMaven(conf, createRequest)

	if err != nil {
		panic(err)
	}

	allRepos, err := GetAll(conf)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, []Repository{Repository{repositoryName, "maven2", "hosted"}}, allRepos)

}
