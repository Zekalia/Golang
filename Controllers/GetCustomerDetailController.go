package Controllers

import (
	"Golang/Managers"
	"Golang/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {
	var output Models.Output
	elektronik, err := Managers.GetElektronikList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}
	buku, err := Managers.GetBukuList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Buku Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}
	mainan, err := Managers.GetMainanList(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Mainan Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}

	if err != nil {
		output.ResponseCode = "99"
		output.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	} else {

		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"
		output.ListElektronik = elektronik
		output.ListBuku = buku
		output.ListMainan = mainan

		c.JSON(http.StatusOK, output)
	}
}

// Call API GetData - Electronik
func GetDataElektronic(c *gin.Context) {
	var outputElectronic Models.OutputElectronic
	elektronik, err := Managers.GetDataElektronic(c)
	if err != nil {
		outputElectronic.ResponseCode = "01"
		outputElectronic.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	}

	if err != nil {
		outputElectronic.ResponseCode = "99"
		outputElectronic.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	} else {
		outputElectronic.ResponseCode = "00"
		outputElectronic.ResponseMessage = "Transaction Success"
		outputElectronic.Elektronik = elektronik
		c.JSON(http.StatusOK, outputElectronic)
	}
}

// Call API Insert Data - Electronik
func GetInsertDataElektronic(c *gin.Context) {
	var outputElectronic Models.OutputElectronic
	_, err := Managers.GetInsertDataElektronic(c)
	if err != nil {
		outputElectronic.ResponseCode = "01"
		outputElectronic.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	}

	if err != nil {
		outputElectronic.ResponseCode = "99"
		outputElectronic.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	} else {
		outputElectronic.ResponseCode = "00"
		outputElectronic.ResponseMessage = "Transaction Success"
		c.JSON(http.StatusOK, outputElectronic)
	}
}

// Call API Update Data - Electronik
func GetUpdateDataElectronic(c *gin.Context) {
	var outputElectronic Models.OutputElectronic
	elektronik, err := Managers.GetUpdateDataElectronic(c)
	if err != nil {
		outputElectronic.ResponseCode = "01"
		outputElectronic.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	}

	if err != nil {
		outputElectronic.ResponseCode = "99"
		outputElectronic.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	} else {
		outputElectronic.ResponseCode = "00"
		outputElectronic.ResponseMessage = "Transaction Success"
		outputElectronic.Elektronik = elektronik
		c.JSON(http.StatusOK, outputElectronic)
	}
}

// Call API Delete Data - Electronik
func DeleteDataElectronic(c *gin.Context) {
	var outputElectronic Models.OutputElectronic
	_, err := Managers.DeleteDataElectronic(c)
	if err != nil {
		outputElectronic.ResponseCode = "01"
		outputElectronic.ResponseMessage = "Elektronik Not Found"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	}

	if err != nil {
		outputElectronic.ResponseCode = "99"
		outputElectronic.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, outputElectronic)
		c.Abort()
	} else {
		outputElectronic.ResponseCode = "00"
		outputElectronic.ResponseMessage = "Transaction Success"
		c.JSON(http.StatusOK, outputElectronic)
	}
}

func GetDetailToy(c *gin.Context) {
	var output Models.OutputToy
	mainan, err := Managers.GetMainanById(c)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Mainan Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}

	if err != nil {
		output.ResponseCode = "99"
		output.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	} else {
		resultMap := map[string]interface{}{
			"mainan": mainan,
		}
		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"
		output.Data = resultMap

		c.JSON(http.StatusOK, output)
	}
}

func UpdateDetailToy(c *gin.Context) {
	var output Models.OutputToy
	_, err := Managers.UpdateMainanById(c, &output)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Mainan Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}

	if err != nil {
		output.ResponseCode = "99"
		output.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	} else {
		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"

		c.JSON(http.StatusOK, output)
	}
}

func DeleteDetailToy(c *gin.Context) {
	var output Models.OutputToy
	_, err := Managers.DeleteMainanById(c, &output)
	if err != nil {
		output.ResponseCode = "01"
		output.ResponseMessage = "Mainan Not Found"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	}

	if err != nil {
		output.ResponseCode = "99"
		output.ResponseMessage = "General Error"
		c.JSON(http.StatusInternalServerError, output)
		c.Abort()
	} else {
		output.ResponseCode = "00"
		output.ResponseMessage = "Transaction Success"

		c.JSON(http.StatusOK, output)
	}
}
