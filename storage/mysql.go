package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ankeshnirala/go/aws-iam-service/logger"
	"github.com/ankeshnirala/go/aws-iam-service/types"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap/zapcore"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	connStr := os.Getenv("MYSQLDB_URL")

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		logger.Fatal(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	logger.Info("Connected to MySQLDB!", zapcore.Field{Type: zapcore.SkipType})

	// defer db.Close()
	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) Find() (*sql.Rows, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer s.db.Close()
	return rows, nil
}

func (s *MySQLStore) FindByID(id int) (*sql.Rows, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return rows, nil
}

func (s *MySQLStore) Registeration(user *types.User) (*sql.Result, error) {

	result, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password, isAdmin) VALUES(?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	defer s.db.Close()
	return &result, nil
}

func (s *MySQLStore) FindByEmail(email string) (*sql.Rows, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		return nil, fmt.Errorf("%s is registered", email)
	}

	defer rows.Close()
	return rows, nil
}
