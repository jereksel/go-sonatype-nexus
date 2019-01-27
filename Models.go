package main

type CreateHostedMavenRepositoryRequest struct {
	Name string `json:"name"`
}

type CreateProxyMavenRepositoryRequest struct {
	Name   string `json:"name"`
	Remote string `json:"remote"`
}
