package src

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Schemes http,https
// @Accept json
// @Produce json
// @Param filename query string true "The filename"
// @Success 200 {object} H{filename string} "Success"
// @Router /createFile [post]
func CreateFile(g *gin.Context) {
	filename := g.Query("filename")

	// Create the file.
	file, _ := os.Create("file/" + filename)
	defer file.Close()

	// Return the filename.
	g.JSON(http.StatusOK, gin.H{"filename": filename})
}
