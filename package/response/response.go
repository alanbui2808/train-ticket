package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	/*
		It tells the Go encoding/json package to map the struct field Code to the JSON key "code" during:

			1. Marshaling (converting Go struct to JSON)
			2. Unmarshaling (converting JSON to Go struct)

		data := ResponseData{Code: 200}
		jsonData, _ := json.Marshal(data)
	*/
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		// msg can be accessed as long as it's in the same package
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		// msg can be accessed as long as it's in the same package
		Message: msg[code],
		Data:    nil,
	})
}
