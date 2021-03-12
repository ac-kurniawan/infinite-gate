package project

import (
	"time"
)

type DataResponse struct {
	ID          int       `json:"id" `
	Name        string    `json:"name"`
	AccessLevel int8      `json:"accessLevel"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SuccessResponseFind struct {
	Code    int          `json:"code"`
	Payload DataResponse `json:"payload"`
}

type SuccessResponseFinds struct {
	Code    int            `json:"code"`
	Payload []DataResponse `json:"payload"`
}

type ErrorResponseProject struct {
	Code    int    `json:"code"`
	Payload string `json:"payload"`
}
