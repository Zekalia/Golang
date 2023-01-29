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
		grp1.GET("getallproduct", Controllers.GetDetail)
		grp1.GET("getDataElectronic", Controllers.GetDataElektronic)             //Call API GetData - Electronik
		grp1.GET("getInsertDataElectronic", Controllers.GetInsertDataElektronic) //Call API Insert Data - Electronik
		grp1.GET("getUpdateDataElectronic", Controllers.GetUpdateDataElectronic) //Call API Update Data - Electronik
		grp1.GET("deleteDataElectronic", Controllers.DeleteDataElectronic)       //Call API Delete Data - Electronik
		grp1.GET("getDetailToy", Controllers.GetDetailToy)
		grp1.POST("UpdateDetailToy", Controllers.UpdateDetailToy)
		grp1.POST("DeleteDetailToy", Controllers.DeleteDetailToy)
	}

	r.Run(":8080")

	return r
}
