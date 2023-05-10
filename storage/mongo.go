package storage

import (
	"context"
	"os"
	"time"

	"github.com/ankeshnirala/go/aws-iam-service/logger"
	"github.com/ankeshnirala/go/aws-iam-service/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap/zapcore"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "users"

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore() (*MongoStore, error) {
	connStr := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(connStr)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err.Error(), zapcore.Field{Type: zapcore.SkipType})

	}

	logger.Info("Connected to MongoDB!", zapcore.Field{Type: zapcore.SkipType})

	return &MongoStore{db: client.Database(DATABASE)}, nil
}

func (s *MongoStore) Find() (*mongo.Cursor, error) {
	return s.db.Collection(COLLECTION).Find(context.TODO(), bson.M{})
}

func (s *MongoStore) FindByID(id string) (*mongo.SingleResult, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.db.Collection(COLLECTION).FindOne(context.TODO(), bson.M{"_id": objectId}), nil
}

func (s *MongoStore) CreateUser(user *types.User) (*mongo.InsertOneResult, error) {
	// s.db.

	insertResult, err := s.db.Collection(COLLECTION).InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (s *MongoStore) GetUserByEmail(email string) *types.User {
	var user *types.User
	s.db.Collection(COLLECTION).FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	return user
}
