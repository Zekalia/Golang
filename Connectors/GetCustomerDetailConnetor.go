package Connectors

import (
	"Golang/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductDetail(c *gin.Context, request *Models.ReqGetBuku) (err error) {
	var reqBuku Models.ReqGetBuku
	reqBuku.Buku = request.Buku

	fmt.Println("reqDetails : ", reqBuku.Buku)

	reqDetailsByte := new(bytes.Buffer)
	json.NewEncoder(reqDetailsByte).Encode(reqBuku)

	var apiUrl = "http://localhost:8080/api/GetBukuByJudulOrcl"
	client := http.Client{}

	req, err := http.NewRequest("GET", apiUrl, reqDetailsByte)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		// "X-Gateway-APIKey": {"a50cfeba-8c04-4a84-aa84-eb0191019606"},
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		var resultResponse Models.Product
		resultMap := map[string]string{}

		response, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(response))
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		json.Unmarshal([]byte(string(response)), &resultMap)
		fmt.Println(resultMap)
		resultResponse.Id = resultMap["Id"]
		resultResponse.Produk = resultMap["Judul buku"]
		resultResponse.Stock = resultMap["Stok"]

		var output Models.Output
		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"
		output.ListProduct = resultResponse

		c.JSON(http.StatusOK, output)
	}
	return
}
