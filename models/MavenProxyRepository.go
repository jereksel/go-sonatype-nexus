package models

type MavenProxyRepository struct {
	Name   string `json:"name"`
	Remote string `json:"remote"`
}

func (r MavenProxyRepository) GetName() string {
	return r.Name
}

func (r MavenProxyRepository) GetFormat() string {
	return "maven"
}

func (r MavenProxyRepository) GetType() string {
	return "proxy"
}
