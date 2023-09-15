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
// @Success 200 {string} string "File deleted successfully"
// @Router /deleteFile [delete]
func deleteFile(g *gin.Context) {
	filename := g.Query("filename")

	os.Remove("file/" + filename)

	// Return success message.
	g.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
