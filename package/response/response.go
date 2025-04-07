package response

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
