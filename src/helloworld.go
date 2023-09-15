package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
