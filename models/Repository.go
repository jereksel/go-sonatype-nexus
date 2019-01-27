package models

type Repository interface {
	GetName() string
	GetFormat() string
	GetType() string
}
