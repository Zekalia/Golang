package Routes

import (
	"Golang/Controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("api/")
	{
		grp1.GET("welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello localhost!")
		})
		grp1.POST("get-detail-product", Controllers.GetDetail)
		grp1.POST("product", Controllers.GetProduct)
	}

	return r
}
