package storage

import (
	"database/sql"
	"os"

	"github.com/ankeshnirala/go/aws-iam-service/types"
	"go.mongodb.org/mongo-driver/mongo"
)

var DATABASE string = os.Getenv("DATABASE_NAME")

type MongoStorage interface {
	CreateUser(*types.User) (*mongo.InsertOneResult, error)
	GetUserByEmail(string) *types.User

	Find() (*mongo.Cursor, error)
	FindByID(string) (*mongo.SingleResult, error)
}

type MySQLStorage interface {
	Find() (*sql.Rows, error)
	FindByID(int) (*sql.Rows, error)
	FindByEmail(string) (*sql.Rows, error)
	Registeration(*types.User) (*sql.Result, error)
}

type RedisStorage interface{}
