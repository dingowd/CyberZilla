package storage

import "github.com/dingowd/CyberZilla/test3/models"

type Storage interface {
	Connect(dsn string) error
	Close() error
	CreateUser(admin, user models.User) error
	ViewUser(admin models.User, user string) (models.User, error)
	UpdateUser(admin, user models.User) error
	DeleteUser(admin models.User, user string) error
}
