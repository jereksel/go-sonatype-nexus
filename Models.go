package main

type CreateHostedMavenRepositoryRequest struct {
	Name string `json:"name"`
}

type CreateProxyMavenRepositoryRequest struct {
	Name   string `json:"name"`
	Remote string `json:"remote"`
}

type CreateGroupMavenRepositoryRequest struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}
