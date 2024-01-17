package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/JWT-Authentication-Golang_MongoDB/Controllers"
)

const (
	WEBPORT = ":8083"
)

func main() {

	router := gin.Default()

	router.GET("/signup", Controllers.SignUp)
	router.Run(WEBPORT)
}
