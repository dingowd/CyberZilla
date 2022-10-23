package mysql

import (
	"github.com/dingowd/CyberZilla/test3/internal/logger"
	"github.com/dingowd/CyberZilla/test3/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB  *sqlx.DB
	Log logger.Logger
}

func New(log logger.Logger) *Storage {
	return &Storage{Log: log}
}

func (s *Storage) Connect(dsn string) error {
	var err error
	s.DB, err = sqlx.Open("mysql", dsn)
	if err == nil {
		s.Log.Info("База " + dsn + " подключена")
	} else {
		s.Log.Error("Ошибка соединения с базой. Проверьте параметры подключения")
	}
	return err
}

func (s *Storage) Close() error {
	s.Log.Info("Закрытие соединения с БД")
	return s.DB.Close()
}

func (s *Storage) CreateUser(admin, user models.User) error {
	return nil
}

func (s *Storage) ViewUser(admin models.User, user string) (models.User, error) {
	var u models.User
	return u, nil
}

func (s *Storage) UpdateUser(admin, user models.User) error {
	return nil
}

func (s *Storage) DeleteUser(admin models.User, user string) error {
	return nil
}
