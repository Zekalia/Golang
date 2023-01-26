package Managers

import (
	"Golang/Connectors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductDetail(c *gin.Context) (err error) {
	var dataMap map[string]any

	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed read request body")
		fmt.Println(err)
		return err
	}
	json.Unmarshal(reqBody, &dataMap)

	id := dataMap["id"].(string)
	if id == "" {
		fmt.Println("Invalid body request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
		return
	}

	fmt.Println("dataMap : ", dataMap)
	fmt.Println("id : ", id)

	err = Connectors.GetProductDetail(c, dataMap)
	if err != nil {
		fmt.Println("Faied get data")
		fmt.Println(err)
		return err
	}
	return
}

func BookDetail(c *gin.Context) (err error) {
	var dataMap map[string]any

	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed read request body")
		fmt.Println(err)
		return err
	}
	json.Unmarshal(reqBody, &dataMap)

	judul := dataMap["buku"].(string)
	if judul == "" {
		fmt.Println("Invalid body request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
		return
	}

	fmt.Println("dataMap : ", dataMap)
	fmt.Println("judul : ", judul)

	err = Connectors.GetBookDetail(c, dataMap)
	if err != nil {
		fmt.Println("Faied get data")
		fmt.Println(err)
		return err
	}
	return
}
