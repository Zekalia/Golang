package Managers

import (
	"fmt"
	"serverLocalGo/Connectors"

	"github.com/gin-gonic/gin"
)

func CustomerDetail(c *gin.Context) (err error) {

	mobileNumber := c.Params.ByName("mobileNumber")
	fmt.Println("mobileNumber Service : ", mobileNumber)
	err = Connectors.GetCustomerDetail(c)
	if err != nil {
		fmt.Println("Faied get data")
		fmt.Println(err)
		return err
	}
	return
}
