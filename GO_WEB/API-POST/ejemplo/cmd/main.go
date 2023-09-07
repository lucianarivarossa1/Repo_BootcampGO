package main

import (
	"gobases/GO_WEB/API-POST/ejemplo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	rt := gin.New()
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())
	db := make([]*handlers.Product, 0)
	ct := handlers.NewControllerProducts(db, 0)
	rt.POST("/products", ct.Save())
	if err := rt.Run(":8084"); err != nil {
		panic(err)
	}
}
