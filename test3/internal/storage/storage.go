package storage

import "github.com/dingowd/CyberZilla/test3/models"

type Storage interface {
	Connect(dsn string) error
	Close() error
	CreateUser(admin string, user models.User) error
	ViewUser(admin string, user string) (models.User, error)
	UpdateUser(admin string, user models.User) error
	DeleteUser(admin string, user string) error
}
