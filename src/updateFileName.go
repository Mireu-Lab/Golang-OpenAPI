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
// @Param newFilename query string true "The new filename"
// @Success 200 {string} string "File name updated successfully"
// @Router /updateFileName [put]
func updateFileName(g *gin.Context) {
	filename := g.Query("filename")
	newFilename := g.Query("newFilename")

	// Check if the file exists.
	if _, err := os.Stat("file/" + filename); os.IsNotExist(err) {
		g.JSON(http.StatusNotFound, H{"error": err.Error()})
		return
	}

	// Rename the file.
	os.Rename("file/"+filename, "file/"+newFilename)

	// Return success message.
	g.JSON(http.StatusOK, H{"message": "File name updated successfully"})
}
