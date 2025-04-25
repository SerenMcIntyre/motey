package main

import (
	"context"
	"log"
	"motey-api/controllers"
	db "motey-api/db"
	"motey-api/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db_test = make(map[string]string)

func setupRouter(queries *db.Queries) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	{
		v1 := r.Group("api/v1")

		registerControllers(v1, queries)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db_test[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db_test[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func registerControllers(v1 *gin.RouterGroup, queries *db.Queries) {
	taskController := controllers.NewTaskController(services.NewTaskService(queries))

	taskController.RegisterRoutes(taskController, v1)
}

func getEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	bg := context.Background()
	log.Println(bg)
	conn, err := pgx.Connect(bg, getEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	r := setupRouter(queries)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
