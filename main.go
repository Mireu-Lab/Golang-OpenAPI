package main

import (
	docs "test/docs"
	"test/src"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r.PUT("/updateFileName", src.updateFileName)
	r.DELETE("/deleteFile", src.deleteFile)
	r.POST("/createFile", src.postfile)
	r.GET("/helloworld", src.Helloworld)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8001")

}
