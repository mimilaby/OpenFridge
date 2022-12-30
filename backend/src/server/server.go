// Package server handles the server
package server

import (
	"context"
	"database/sql"
	"fmt"
	"openfridge/src/config"
	"openfridge/src/controllers"
	dbConn "openfridge/src/db/sqlc"
	"openfridge/src/routes"

	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// pq is the postgres driver788
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server      *gin.Engine
	db          *dbConn.Queries
	mongoClient *mongo.Client

	// AuthController holds the auth controller
	AuthController controllers.AuthController
	// AuthRoutes holds the auth routes
	AuthRoutes routes.AuthRoutes

	// FoodController holds the food controller
	FoodController controllers.FoodController
	// FoodRoutes holds the food routes
	FoodRoutes routes.FoodRoutes
)

// Init initializes the server
func Init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db = ConnectPostgres(config.PostgreDriver, config.PostgresSource)
	mongoClient = ConnectMongo(config.MongoURL, config.MongoUsername, config.MongoPassword)

	AuthController = *controllers.NewAuthController(db)
	AuthRoutes = routes.NewAuthRoutes(AuthController)

	FoodController = *controllers.NewFoodController(db)
	FoodRoutes = routes.NewFoodRoutes(FoodController)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}

	server = gin.Default()
	server.Use(cors.New(corsConfig))
	server.SetTrustedProxies(nil)
}

// Run runs the server
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
	FoodRoutes.FoodRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.Port))
}

// ConnectPostgres connect to postgresql
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

// ConnectMongo connect to MongoDB
func ConnectMongo(mongoURL string, mongoUsername string, mongoPassword string) *mongo.Client {
	// Authentication
	credentials := options.Credential{
		AuthSource: "admin",
		Username:   mongoUsername,
		Password:   mongoPassword,
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURL).
		SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
