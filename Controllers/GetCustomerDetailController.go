package Controllers

import (
	"Golang/Managers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {

	err := Managers.CustomerDetail(c)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}

func GetStok(c *gin.Context) {

	err := Managers.StokList(c)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}
