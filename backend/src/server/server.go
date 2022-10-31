package server

import (
	"database/sql"
	"fmt"
	"homeapp/src/config"
	"homeapp/src/controllers"
	dbConn "homeapp/src/db/sqlc"
	"homeapp/src/query"
	"homeapp/src/routes"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	db          *dbConn.Queries
	mongoClient *mongo.Client

	AuthController controllers.AuthController
	AuthRoutes     routes.AuthRoutes
)

func Init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db = ConnectPostgres(config.PostgreDriver, config.PostgresSource)
	mongoClient = query.ConnectMongo(config.MongoUsername, config.MongoPassword)
	// fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db)
	AuthRoutes = routes.NewAuthRoutes(AuthController)

	server = gin.Default()
	server.SetTrustedProxies(nil)
}

func Run() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.Port))
}

// connect to postgresql
func ConnectPostgres(PostgresDriver string, PostgresSource string) *dbConn.Queries {

	// Set client options
	conn, err := sql.Open(PostgresDriver, PostgresSource)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	}
	db = dbConn.New(conn)
	fmt.Println("Connected to PostgreSQL!")
	return db
}
