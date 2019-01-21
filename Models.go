package main

type BaseData struct {
	Name   string
	Format string
}

type IRepository interface {
	GetName() string
	GetFormat() string
	GetType() string
}

type MavenRepository struct {
	Name   string `json:"name"`
	Format string `json:"format"`
}

func (r MavenRepository) GetName() string {
	return r.Name
}

func (r MavenRepository) GetFormat() string {
	return r.Format
}

func (r MavenRepository) GetType() string {
	return "maven"
}

type CreateHostedMavenRepositoryRequest struct {
	Name string `json:"name"`
}
