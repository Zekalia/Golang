package Connectors

import (
	"Golang/Models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokList(c *gin.Context) (err error) {

	result := make([]Models.RequestGetDetails, 0)
	products := getElektronik()
	for _, product := range products {
		result = append(result, product)
	}

	products = getBuku()
	for _, product := range products {
		result = append(result, product)
	}

	products = getMinuman()
	for _, product := range products {
		result = append(result, product)
	}

	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, string(response))

	return
}

func getElektronik() (products []Models.RequestGetDetails) {
	response, err := ioutil.ReadFile("daftarelektronik.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(response, &products)

	if err != nil {
		panic(err)
	}

	return products
}

func getBuku() (products []Models.RequestGetDetails) {
	response, err := ioutil.ReadFile("daftarbuku.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(response, &products)

	if err != nil {
		panic(err)
	}

	return products
}

func getMinuman() (products []Models.RequestGetDetails) {
	response, err := ioutil.ReadFile("daftarminuman.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(response, &products)

	if err != nil {
		panic(err)
	}

	return products
}

func GetCustomerDetail(c *gin.Context) (err error) {
	var reqDetails Models.RequestGetDetails

	reqDetails.Id = ""
	reqDetails.Produk = ""
	reqDetails.Stock = ""

	return
}
