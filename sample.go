package main

import (
	"net/http"
	"os"

	docs "test/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type H map[string]interface{}

// @BasePath /
// @Schemes http,https
// @Accept json
// @Produce json
// @SecurityDefinitions.basic BasicAuth
// @Security "BasicAuth"
// @Router /protected [get]
func protected(c *gin.Context) {
	// Get the username and password from the Authorization header.
	username, password, _ := c.Request.BasicAuth()

	// Check if the username and password are valid.
	if username != "admin" || password != "secret" {
		c.JSON(http.StatusUnauthorized, H{"error": "Invalid username or password"})
		return
	}

	// Return a success message.
	c.JSON(http.StatusOK, H{"message": "You are authorized"})
}

// @BasePath /
// @Schemes http,https
// @Accept json
// @Produce json
// @Param filename query string true "The filename"
// @Param newFilename query string true "The new filename"
// @Success 200 {string} string "File name updated successfully"
// @Router /updateFileName [put]
func updateFileName(c *gin.Context) {
	filename := c.Query("filename")
	newFilename := c.Query("newFilename")

	// Check if the file exists.
	if _, err := os.Stat("file/" + filename); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Rename the file.
	os.Rename("file/"+filename, "file/"+newFilename)

	// Return success message.
	c.JSON(http.StatusOK, gin.H{"message": "File name updated successfully"})
}

// @BasePath /
// @Schemes http,https
// @Accept json
// @Produce json
// @Param filename query string true "The filename"
// @Success 200 {string} string "File deleted successfully"
// @Router /deleteFile [delete]
func deleteFile(c *gin.Context) {
	filename := c.Query("filename")

	os.Remove("file/" + filename)

	// Return success message.
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// @BasePath /
// @Schemes http,https
// @Accept json
// @Produce json
// @Param filename query string true "The filename"
// @Success 200 {object} H{filename string} "Success"
// @Router /createFile [post]
func postfile(g *gin.Context) {
	filename := g.Query("filename")

	// Create the file.
	file, _ := os.Create("file/" + filename)
	defer file.Close()

	// Return the filename.
	g.JSON(http.StatusOK, gin.H{"filename": filename})
}

// @BasePath /
// PingExample godoc
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSONP(http.StatusOK, H{"msg": "helloworld"})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// r.PATCH()
	r.GET("/protected", protected)

	r.PUT("/updateFileName", updateFileName)
	r.DELETE("/deleteFile", deleteFile)
	r.POST("/createFile", postfile)
	r.GET("/helloworld", Helloworld)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":80")

}
