package api

import (
	"github.com/ankeshnirala/go/aws-iam-service/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	mongoStore storage.MongoStorage
	mysqlStore storage.MySQLStorage
	redisStore storage.RedisStorage
}

func NewServer(listenAddr string, mongoStore storage.MongoStorage, mysqlStore storage.MySQLStorage, redisStore storage.RedisStorage) *Server {
	return &Server{
		listenAddr: listenAddr,
		mongoStore: mongoStore,
		mysqlStore: mysqlStore,
		redisStore: redisStore,
	}
}

func (s *Server) Start() error {
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// authenticated routes
	protectedRoute := router.Group("/users")
	protectedRoute.Use(Authenticate())

	// auth routes
	router.POST("/signup", s.handleSignup)
	router.POST("/login", s.handleLogin)

	protectedRoute.POST("/add", s.handleMemberSignup)

	protectedRoute.GET("/", s.handleGetUsers)
	protectedRoute.GET("/profile", s.handleGetUserProfile)

	return router.Run(s.listenAddr)
}
