package main

type MavenHostedRepository struct {
	Name string `json:"name"`
}

func (r MavenHostedRepository) GetName() string {
	return r.Name
}

func (r MavenHostedRepository) GetFormat() string {
	return "maven"
}

func (r MavenHostedRepository) GetType() string {
	return "hosted"
}
