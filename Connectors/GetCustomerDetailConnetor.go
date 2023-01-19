package Connectors

import (
	"github.com/gin-gonic/gin"
	"serverLocalGo/Models"
)

func GetCustomerDetail(c *gin.Context) (err error) {
	var reqDetails Models.RequestGetDetails

	reqDetails.Id = ""
	reqDetails.Produk = ""
	reqDetails.Stock = ""

	return
}
