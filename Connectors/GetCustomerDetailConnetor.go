package Connectors

import (
	"Golang/Models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokList(c *gin.Context) (err error) {

	resultElektronik := make([]Models.RequestGetDetails, 0)
	resultBuku := make([]Models.RequestGetDetails, 0)
	resultMinuman := make([]Models.RequestGetDetails, 0)

	getElektronik, err := ioutil.ReadFile("daftarelektronik.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(getElektronik, &resultElektronik)

	getBuku, err := ioutil.ReadFile("daftarbuku.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(getBuku, &resultBuku)

	getMinuman, err := ioutil.ReadFile("daftarminuman.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(getMinuman, &resultMinuman)

	var output Models.Output
	output.ResponseCode = "00"
	output.ResponseMessage = "Transaction Success"
	output.ListElektronik = resultElektronik
	output.ListBuku = resultBuku
	output.ListMinuman = resultMinuman

	c.JSON(http.StatusOK, output)

	return
}

func GetCustomerDetail(c *gin.Context) (err error) {
	var reqDetails Models.RequestGetDetails

	reqDetails.Id = ""
	reqDetails.Produk = ""
	reqDetails.Stock = ""

	return
}
