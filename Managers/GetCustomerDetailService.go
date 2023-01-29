package Managers

import (
	"Golang/Connectors"
	"Golang/Models"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func GetElektronikList(c *gin.Context) (Elektronik []Models.Product, err error) {
	elektronik, err := Connectors.GetElektronikList(c)
	if err != nil {
		return nil, err
	}
	return elektronik, nil
}

func GetBukuList(c *gin.Context) (Buku []Models.Product, err error) {
	buku, err := Connectors.GetBukuList(c)
	if err != nil {
		return nil, err
	}
	return buku, nil
}

func GetMainanList(c *gin.Context) (Mainan []Models.Product, err error) {
	mainan, err := Connectors.GetMainanList(c)
	if err != nil {
		return nil, err
	}
	return mainan, nil
}

// Call API GetData - Electronik
func GetDataElektronic(c *gin.Context) (Elektronik Models.Product, err error) {
	elektronik, err := Connectors.GetDataElektronic(c)
	if err != nil {
		return elektronik, err
	}
	return elektronik, nil
}

// Call API Insert Data - Electronik
func GetInsertDataElektronic(c *gin.Context) (Elektronik Models.Product, err error) {
	elektronik, err := Connectors.GetInsertDataElektronic(c)
	if err != nil {
		return elektronik, err
	}
	return elektronik, nil
}

// Call API Update Data - Electronik
func GetUpdateDataElectronic(c *gin.Context) (Elektronik Models.Product, err error) {
	elektronik, err := Connectors.GetUpdateDataElectronic(c)
	if err != nil {
		return elektronik, err
	}
	return elektronik, nil
}

// Call API Delete Data - Electronik
func DeleteDataElectronic(c *gin.Context) (Elektronik Models.Product, err error) {
	elektronik, err := Connectors.DeleteDataElectronic(c)
	if err != nil {
		return elektronik, err
	}
	return elektronik, nil
}

func GetMainanById(c *gin.Context) (Mainan Models.Product, err error) {
	var dataMap map[string]interface{}

	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed read request body")
		fmt.Println(err)
		return
	}
	json.Unmarshal(reqBody, &dataMap)

	mainan, err := Connectors.GetMainanById(c, dataMap)
	if err != nil {
		return mainan, err
	}
	return mainan, nil
}

func UpdateMainanById(c *gin.Context, output *Models.OutputToy) (Mainan Models.Product, err error) {
	var dataMap map[string]interface{}

	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed read request body")
		fmt.Println(err)
		return
	}
	json.Unmarshal(reqBody, &dataMap)

	mainan, err := Connectors.UpdateMainanById(c, dataMap, output)
	if err != nil {
		return mainan, err
	}
	return mainan, nil
}

func DeleteMainanById(c *gin.Context, output *Models.OutputToy) (Mainan Models.Product, err error) {
	var dataMap map[string]interface{}

	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed read request body")
		fmt.Println(err)
		return
	}
	json.Unmarshal(reqBody, &dataMap)

	mainan, err := Connectors.DeleteMainanById(c, dataMap, output)
	if err != nil {
		return mainan, err
	}
	return mainan, nil
}
