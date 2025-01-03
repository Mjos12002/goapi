package customer

import "advanced/model"

type ResponseInfo struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    model.Customer `json:"data"`
}
