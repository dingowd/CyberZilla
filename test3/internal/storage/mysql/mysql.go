package mysql

import (
	"errors"
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

func (s *Storage) CreateUser(admin string, user models.User) error {
	if !s.IsAdmin(admin) {
		return errors.New("User " + admin + " is not admin")
	}
	if s.IsUserExist(user.Name) {
		return errors.New("User " + user.Name + " already exists")
	}
	query := `insert into users(name, group_id, pass, email) values(?, ?, ?, ?)`
	if _, err := s.DB.Exec(query, user.Name, user.GroupId, user.Pass, user.Email); err != nil {
		return err
	}
	return nil
}

func (s *Storage) ViewUser(admin string, user string) (models.User, error) {
	var u models.User
	if !s.IsAdmin(admin) {
		return u, errors.New("User " + admin + " is not admin")
	}
	query := `select * from users where name = ?`
	row := s.DB.QueryRow(query, user)
	if err := row.Scan(&u.ID, &u.Name, &u.GroupId, &u.Pass, &u.Email); err != nil {
		return u, err
	}
	return u, nil
}

func (s *Storage) UpdateUser(admin string, user models.User) error {
	if !s.IsAdmin(admin) {
		return errors.New("User " + admin + " is not admin")
	}
	if !s.IsUserExist(user.Name) {
		return errors.New("User " + user.Name + " is not exists")
	}
	query := `update users set group_id = ?, pass = ?, email = ? where name = ?`
	if _, err := s.DB.Exec(query, user.GroupId, user.Pass, user.Email, user.Name); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteUser(admin string, user string) error {
	if admin == user {
		return errors.New("Unable to delete yourself")
	}
	if !s.IsAdmin(admin) {
		return errors.New("User " + admin + " is not admin")
	}
	if !s.IsUserExist(user) {
		return errors.New("User " + user + " is not exists")
	}
	query := `delete from users where name = ?`
	if _, err := s.DB.Exec(query, user); err != nil {
		return err
	}
	return nil
}

func (s *Storage) IsAdmin(admin string) bool {
	var r models.Rights
	query := `  select users.name, r_create from users 
				inner join usersgroups on users.group_id = usersgroups.group_id 
				where users.name = ?`
	row := s.DB.QueryRow(query, admin)
	if err := row.Scan(&r.Name, &r.Right); err != nil {
		return false
	}
	if !r.Right {
		return false
	}
	return true
}

func (s *Storage) IsUserExist(name string) bool {
	var n string
	query := `select name from users where name = ?`
	row := s.DB.QueryRow(query, name)
	if err := row.Scan(&n); err != nil {
		return false
	}
	return true
}
