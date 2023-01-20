package Connectors

import (
	"Golang/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductDetail(c *gin.Context, dataMap map[string]any) (err error) {
	//var reqDetail Models.Product
	var resDetail Models.Product

	//reqDetail.Id = dataMap["id"].(string)

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(dataMap)

	var apiUrl = "http://localhost:8080/api/product"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqBody)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	// req.Header = http.Header{
	// 	"Content-Type":     {"application/json"},
	// 	"X-Gateway-APIKey": {"a50cfeba-8c04-4a84-aa84-eb0191019606"},
	// }

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		response, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		err = json.Unmarshal(response, &resDetail)
		if err != nil {
			log.Fatalf("failed to convert response body: %s", err)
		}
		c.JSON(http.StatusOK, resDetail)
	}

	return
}
