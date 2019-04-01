package router

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/fuyi-huang/backend/utils"
	"github.com/gin-gonic/gin"
)

func process(params url.Values) (int, string) {
	necces := []string{"appid", "type", "time", "sign"}
	if !utils.CheckParam(necces, params) {
		return 1001, "Lack of necessary parameters"
	}
	if !utils.ValidateSign(params) {
		return 1002, "wrong sign"
	}
	type_, err := strconv.Atoi(params.Get("type"))
	if err != nil {
		return 1001, "wrong parameters"
	}
	if type_ == 1 {
		order_number := utils.OrderNumber()
		return 0, fmt.Sprintf("%d", order_number)
	}
	return 0, utils.OrderStr()
}

// GET /get_order
func GetOrder(c *gin.Context) {
	params := c.Request.URL.Query()
	error_code, res := process(params)
	status := 1
	if error_code != 0 {
		status = 0
	}

	c.JSON(200, gin.H{
		"status":     status,
		"error_code": error_code,
		"data":       res,
	})
}
