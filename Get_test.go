package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetNotExisting(t *testing.T) {

	//Given
	conf := getConf()
	removeAll(conf)

	//When
	notExisting, err := Get(conf, "ABC")
	if err != nil {
		panic(err)
	}

	//Then
	assert.Nil(t, notExisting)

}

func TestMavenHosted(t *testing.T) {

	//Given
	conf := getConf()

	removeAll(conf)

	hostedMavenName := "a-hosted-maven"

	if err := CreateHostedMaven(conf, CreateHostedMavenRepositoryRequest{hostedMavenName}); err != nil {
		panic(err)
	}

	//When
	repo, err := Get(conf, hostedMavenName)
	if err != nil {
		panic(err)
	}

	//Then
	assert.Equal(t,
		MavenHostedRepository{hostedMavenName},
		*repo,
	)

}
