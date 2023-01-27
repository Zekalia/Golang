package Managers

import (
	"Golang/Connectors"
	"Golang/Models"

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
