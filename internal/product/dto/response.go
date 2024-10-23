package dto

type ResponseModels struct {
	ResponseCode    int         `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	Data            interface{} `json:"data"` // Use interface{} to support various types of data
}
