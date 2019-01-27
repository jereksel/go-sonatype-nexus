package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestRemoveExistingRepository(t *testing.T) {

	conf := getConf()

	removeAll(conf)

	repository1Name := uuid.New().String()
	repository2Name := uuid.New().String()

	{
		createRequest := CreateHostedMavenRepositoryRequest{repository1Name}

		if err := CreateHostedMaven(conf, createRequest); err != nil {
			panic(err)
		}

	}

	{
		createRequest := CreateHostedMavenRepositoryRequest{repository2Name}

		if err := CreateHostedMaven(conf, createRequest); err != nil {
			panic(err)
		}

	}

	if err := Remove(conf, repository1Name); err != nil {
		panic(err)
	}

	allRepos, err := GetAll(conf)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, []Repository{Repository{repository2Name, "maven2", "hosted"}}, allRepos)

}
