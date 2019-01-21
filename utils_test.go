package main

import "os"

func getNexusHostname() string {
	if len(os.Getenv("CI")) > 0 {
		return os.Getenv("NEXUS_HOSTNAME")
	} else {
		return "localhost"
	}
}

func getConf() Configuration {
	hostname := getNexusHostname()
	return Configuration{
		"http://" + hostname + ":8081",
		"admin",
		"admin123",
	}
}

func removeAll(conf Configuration) {

	allRepos, err := GetAll(conf)

	if err != nil {
		panic(err)
	}

	for _, repo := range allRepos {
		if err := Remove(conf, repo.Name); err != nil {
			panic(err)
		}
	}

}
