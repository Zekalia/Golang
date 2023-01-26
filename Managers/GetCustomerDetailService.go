package Managers

import (
	"Golang/Connectors"
	"Golang/Models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProductDetail(c *gin.Context, request *Models.ReqGetBuku) (err error) {
	err = Connectors.GetProductDetail(c, request)
	if err != nil {
		fmt.Println("Faied get data")
		fmt.Println(err)
		return err
	}
	return
}
