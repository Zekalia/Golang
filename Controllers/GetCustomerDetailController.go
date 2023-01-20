package Controllers

import (
	"Golang/Managers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {

	err := Managers.ProductDetail(c)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}

func GetProduct(c *gin.Context) {
	type product struct {
		Id string `json:"id"`
	}

	var reqBody product

	c.BindJSON(&reqBody)

	c.JSON(http.StatusOK, gin.H{"id": reqBody.Id, "produk": "buku", "stock": "10"})

}
