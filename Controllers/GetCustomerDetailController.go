package Controllers

import (
	"Golang/Managers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {

	err := Managers.ProductDetail(c)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}
