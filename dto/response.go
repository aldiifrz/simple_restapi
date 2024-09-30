package dto

import "simple_restapi/entity"

type ResponseModels struct {
	ResponseCode    int              `json:"response_code"`
	ResponseMessage string           `json:"response_message"`
	Data            []entity.Product `json:"data"`
}
