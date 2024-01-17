package Controllers

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SignUp(c *gin.Context) {

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// 	log.Println(err)
	// }
	_ = godotenv.Load()
	uri := os.Getenv("MONGODB_URI")
	if len(uri) == 0 {
		println("No .env file found")
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		return
	}
	println(uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error Connecting: ", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error Connecting: ", err)
		}
	}()

	c.JSON(200, gin.H{
		"client": client,
	})
}
