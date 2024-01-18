package Controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/JWT-Authentication-Golang_MongoDB/Models"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SignUp(c *gin.Context) {

	// read the body
	var jsonData Models.User_SignUp
	// Bind JSON data from the request body
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

	coll := client.Database("Jwt-Golang").Collection("Users")

	email_filter := bson.D{{"email", jsonData.Email}}
	username_filter := bson.D{{"username", jsonData.Username}}

	var result Models.User_SignUp
	err = coll.FindOne(context.TODO(), email_filter).Decode(&result)

	// check if any doc matches with the email

	if err == mongo.ErrNoDocuments {

		// ---- USERNAME----
		err = coll.FindOne(context.TODO(), username_filter).Decode(&result)

		if err == mongo.ErrNoDocuments {
			jsonData.Id = primitive.NewObjectID()

			if jsonData.Password == "" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Password Empty",
				})
				return
			}
			_, err = coll.InsertOne(context.TODO(), jsonData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": string(err.Error()),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Status": "User SignUp Successful",
			})
			return
		}

		// ---- USERNAME----

	}

	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Email or username Already exists",
		})
		return
	}

	fmt.Println(result)
	c.JSON(http.StatusBadRequest, gin.H{
		"client": client,
	})
}
