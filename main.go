package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ankeshnirala/go/aws-iam-service/api"
	"github.com/ankeshnirala/go/aws-iam-service/logger"
	"github.com/ankeshnirala/go/aws-iam-service/storage"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

func main() {

	err := godotenv.Load("app.env")
	if err != nil {
		logger.Error(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	appPort := os.Getenv("APP_PORT")

	listenAddr := flag.String("listenaddr", appPort, "the server address")
	flag.Parse()

	mongoStore, err := storage.NewMongoStore()
	if err != nil {
		logger.Error(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	mysqlStore, err := storage.NewMySQLStore()
	if err != nil {
		logger.Error(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	redisStore, err := storage.NewRedisStore()
	if err != nil {
		logger.Error(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	server := api.NewServer(*listenAddr, mongoStore, mysqlStore, redisStore)
	msg := fmt.Sprintf("Server is running on port %s", *listenAddr)
	logger.Info(msg, zapcore.Field{Type: zapcore.SkipType})
	log.Fatal(server.Start())
}
