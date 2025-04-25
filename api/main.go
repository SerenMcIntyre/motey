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
	"google.golang.org/api/idtoken"
)

var db_test = make(map[string]string)

func setupRouter(queries *db.Queries) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	{
		v1 := r.Group("api/v1")

		registerControllers(v1, queries)
		v1.POST("/login", googleAuthHandler)
	}

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

func googleAuthHandler(c *gin.Context) {
    var body struct {
        IDToken string `json:"id_token"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
        return
    }

    payload, err := idtoken.Validate(context.Background(), body.IDToken, getEnv("GOOGLE_CLIENT_ID"))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        return
    }

    email := payload.Claims["email"].(string)
    sub := payload.Subject

    c.JSON(200, gin.H{"email": email, "user_id": sub})
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

	r.Run(":8080")
}
