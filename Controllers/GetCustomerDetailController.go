package Controllers

import (
	"Golang/Managers"
	"Golang/Models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {
	var reqGetBuku Models.ReqGetBuku

	c.BindJSON(&reqGetBuku)

	err := Managers.ProductDetail(c, &reqGetBuku)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}
