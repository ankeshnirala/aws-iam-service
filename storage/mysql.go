package storage

import (
	"database/sql"
	"os"

	"github.com/ankeshnirala/go/aws-iam-service/logger"
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

	defer db.Close()
	return &MySQLStore{db: db}, nil
}
