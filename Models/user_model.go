package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User_SignUp struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Username string             `bson:"username,omitempty"`
}
