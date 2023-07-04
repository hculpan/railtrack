package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hculpan/railtrack/handlers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to DB")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.IndexHandler)
	r.GET("/api/about", handlers.AboutHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
