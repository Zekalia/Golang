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

	output := map[string]interface{}{
		"response_code":    "00",
		"response_message": "Transaction Success",
		"list_elektronik":  []interface{}{resultElektronik},
		"list_buku":        []interface{}{resultBuku},
		"list_minuman":     []interface{}{resultMinuman},
	}

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
