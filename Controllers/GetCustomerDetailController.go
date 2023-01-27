package Controllers

import (
	"Golang/Managers"
	"Golang/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {
	var output Models.Output
	elektronik, err := Managers.GetElektronikList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}
	buku, err := Managers.GetBukuList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Buku Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}
	mainan, err := Managers.GetMainanList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Mainan Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}

	if err != nil {
		output.ResponseCode = "99"
		output.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	} else {

		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"
		output.ListElektronik = elektronik
		output.ListBuku = buku
		output.ListMainan = mainan

		c.JSON(http.StatusOK, output)
	}
}
