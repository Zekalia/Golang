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
		grp1.POST("getdetail/:mobileNumber", Controllers.GetDetail)
	}

	grp2 := r.Group("api/")
	{
		grp2.POST("getstok/", Controllers.GetStok)
	}

	r.Run(":8888")

	return r
}
