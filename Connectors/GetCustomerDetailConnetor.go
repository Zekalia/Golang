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

// Call API GetData - Electronik
func GetDataElektronic(c *gin.Context) (Elektronik Models.Product, err error) {
	var resultElektronik Models.Product
	var reqBodyGetDataElec Models.ReqGetElectronicById

	reqBody := new(bytes.Buffer)
	c.BindJSON(&reqBodyGetDataElec)

	json.NewEncoder(reqBody).Encode(&reqBodyGetDataElec)

	var apiUrl = "http://localhost:8484/GetData"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqBody)
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
		fmt.Println(resultElektronik)
	}
	return resultElektronik, nil
}

// Call API Insert Data - Electronik
func GetInsertDataElektronic(c *gin.Context) (Elektronik Models.Product, err error) {
	var resultElektronik Models.Product
	var reqBodyInsertDataElec Models.ReqInsertDataElectronic

	reqBody := new(bytes.Buffer)
	c.BindJSON(&reqBodyInsertDataElec)

	json.NewEncoder(reqBody).Encode(&reqBodyInsertDataElec)

	var apiUrl = "http://localhost:8484/InsertData"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqBody)
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

// Call API Update Data - Electronik
func GetUpdateDataElectronic(c *gin.Context) (Elektronik Models.Product, err error) {
	var resultElektronik Models.Product
	var reqBodyUpdateDataElec Models.ReqUpdateDataElectronic

	reqBody := new(bytes.Buffer)
	c.BindJSON(&reqBodyUpdateDataElec)

	json.NewEncoder(reqBody).Encode(&reqBodyUpdateDataElec)

	var apiUrl = "http://localhost:8484/UpdateData"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqBody)
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

// Call API Delete Data - Electronik
func DeleteDataElectronic(c *gin.Context) (Elektronik Models.Product, err error) {
	var resultElektronik Models.Product
	var reqBodyDeleteDataElec Models.ReqGetElectronicById

	reqBody := new(bytes.Buffer)
	c.BindJSON(&reqBodyDeleteDataElec)

	json.NewEncoder(reqBody).Encode(&reqBodyDeleteDataElec)

	var apiUrl = "http://localhost:8484/DeleteData"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqBody)
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
