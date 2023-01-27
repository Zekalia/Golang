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

func GetElektronikList(c *gin.Context) (Elektronik []Models.Product, err error) {
	var resultElektronik []Models.Product

	reqDetailsByte := new(bytes.Buffer)

	var apiUrl = "http://localhost:8181/GetAllElektronik"
	client := http.Client{}

	req, err := http.NewRequest("GET", apiUrl, reqDetailsByte)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	req.Header = http.Header{"Content-Type": {"application/json"}}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		json.Unmarshal([]byte(string(response)), &resultElektronik)
	}
	return resultElektronik, nil
}

func GetBukuList(c *gin.Context) (Buku []Models.Product, err error) {
	var resultBuku []Models.Product

	reqDetailsByte := new(bytes.Buffer)

	var apiUrl = "http://localhost:8282/api/GetAllBukuOrcl"
	client := http.Client{}

	req, err := http.NewRequest("GET", apiUrl, reqDetailsByte)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	req.Header = http.Header{"Content-Type": {"application/json"}}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		json.Unmarshal([]byte(string(response)), &resultBuku)
	}
	return resultBuku, nil
}

func GetMainanList(c *gin.Context) (Mainan []Models.Product, err error) {
	var resultMainan []Models.Product

	reqDetailsByte := new(bytes.Buffer)

	var apiUrl = "http://localhost:8383/api/getAllToys"
	client := http.Client{}

	req, err := http.NewRequest("GET", apiUrl, reqDetailsByte)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	req.Header = http.Header{"Content-Type": {"application/json"}}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		json.Unmarshal([]byte(string(response)), &resultMainan)
	}
	return resultMainan, nil
}
