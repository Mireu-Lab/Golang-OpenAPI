package main

import (
	docs "github.com/Mireu-Lab/Golang-OpenAPI/docs"

	"github.com/Mireu-Lab/Golang-OpenAPI/src"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r.PUT("/updateFileName", src.UpdateFileName)
	r.DELETE("/deleteFile", src.DeleteFile)
	r.POST("/createFile", src.CreateFile)
	r.GET("/helloworld", src.HelloWorld)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8001")

}
