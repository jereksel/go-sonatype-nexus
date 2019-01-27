package main

type MavenGroupRepository struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func (r MavenGroupRepository) GetName() string {
	return r.Name
}

func (r MavenGroupRepository) GetFormat() string {
	return "maven"
}

func (r MavenGroupRepository) GetType() string {
	return "hosted"
}
