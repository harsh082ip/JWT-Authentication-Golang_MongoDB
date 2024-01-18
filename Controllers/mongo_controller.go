package Controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	Helpers "github.com/harsh082ip/JWT-Authentication-Golang_MongoDB/Helpers"
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

	// Load the Environment Variables
	_ = godotenv.Load()

	// Get the Connection String
	uri := os.Getenv("MONGODB_URI")

	// If uri is empty
	if len(uri) == 0 {
		println("No .env file found")
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		return
	}
	println(uri)

	// Trying to Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// If any Error Connecting
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Error Connecting to Database",
			"Error":  err,
		})
		log.Fatal("Error Connecting: ", err)
	}
	defer func() {
		// Release all Resources
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error Connecting: ", err)
		}
	}()

	// Querying the collection
	coll := client.Database("Jwt-Golang").Collection("Users")

	// creating both email and username filters
	email_filter := bson.D{{"email", jsonData.Email}}
	username_filter := bson.D{{"username", jsonData.Username}}

	var result Models.User_SignUp
	err = coll.FindOne(context.TODO(), email_filter).Decode(&result)

	// check if any doc matches with the email

	// checking if the database has a doc with email_filter
	if err == mongo.ErrNoDocuments {

		// Now we'll check same for the username

		// ---- USERNAME----
		err = coll.FindOne(context.TODO(), username_filter).Decode(&result)

		// checking if the database has a doc with username_filter
		if err == mongo.ErrNoDocuments {

			// now we'll try to sign up the user

			jsonData.Id = primitive.NewObjectID()

			// checking if the password is empty
			if jsonData.Password == "" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Password Empty",
				})
				return
			}

			// Checking if the password meets the criteria
			if !Helpers.CheckPasswordValidity(jsonData.Password) {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error":  "password constraints not fulfilled",
					"detail": "Password does not meet the required criteria: it must be at least 8 characters long, contain at least one lowercase letter, one uppercase letter, one special character, and one numeric digit.",
				})
				return
			}

			// Creating Hash of the Password
			hashpass, er := Helpers.HashPassword(jsonData.Password)
			if er != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error":  "Error creating password hash",
					"detail": er,
				})
			}

			// overriding the previous text-pass with hash password
			jsonData.Password = hashpass

			// Attempting to create user in the database
			_, err = coll.InsertOne(context.TODO(), jsonData)

			// handle if any error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": string(err.Error()),
				})
				return
			}

			// User signUp Successful
			c.JSON(http.StatusOK, gin.H{
				"Status": "User SignUp Successful",
			})
			return
		}

		// ---- USERNAME----

	}

	// if err value is nil that means email or password already exists
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Email or username Already exists",
		})
		return
	}

	// If somehow user reaches here
	fmt.Println(result)
	c.JSON(http.StatusBadRequest, gin.H{
		"client": client,
	})
}
