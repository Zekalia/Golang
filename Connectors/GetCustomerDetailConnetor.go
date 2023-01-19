package Connectors

import (
	"Golang/Models"

	"github.com/gin-gonic/gin"
)

func GetCustomerDetail(c *gin.Context) (err error) {
	var reqDetails Models.RequestGetDetails

	reqDetails.Id = ""
	reqDetails.Produk = ""
	reqDetails.Stock = ""

	return
}
