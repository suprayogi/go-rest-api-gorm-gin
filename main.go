package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productcontroller "github.com/suprayogi/go-rest-api/controllers"
	"github.com/suprayogi/go-rest-api/models"
)

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Welcome",
	})
}

func main() {
	r := gin.Default()
	models.ConnectDB()
	r.GET("/", welcome)
	r.GET("/products", productcontroller.Index)
	r.POST("/products", productcontroller.Create)
	r.PUT("/products/:id", productcontroller.Update)
	r.DELETE("/products/:id", productcontroller.Delete)

	r.Run(":8899")
}
