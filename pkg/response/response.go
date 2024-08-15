package response

import (
	"math"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	StatusCode int    `json:"statusCode"`
	TaskStatus bool   `json:"taskStatus"`
	Message    string `json:"message"`
}

type ResponseData struct {
	StatusCode int         `json:"statusCode"`
	TaskStatus bool        `json:"taskStatus"`
	Data       interface{} `json:"data"`
	Pagin      *Pagination `json:"pagin,omitempty"`
}

type Pagination struct {
	PageNumber  int `json:"pageNumber"`
	PageSize    int `json:"pageSize"`
	TotalPages  int `json:"totalPages"`
	TotalRecord int `json:"totalRecord"`
}

func Message(ctx *gin.Context, statusCode int, taskStatus bool, message string) {
	response := ResponseMessage{
		StatusCode: statusCode,
		TaskStatus: taskStatus,
		Message:    message,
	}
	ctx.JSON(statusCode, response)
}

func SendData(ctx *gin.Context, statusCode int, taskStatus bool, data interface{}, pagin *Pagination) {
	response := ResponseData{
		StatusCode: statusCode,
		TaskStatus: taskStatus,
		Data:       data,
		Pagin:      pagin,
	}
	if pagin != nil {
		pagin.TotalPages = calculateTotalPages(pagin.TotalRecord, pagin.PageSize)
		response.Pagin = pagin
	}
	ctx.JSON(statusCode, response)
}

func calculateTotalPages(totalRecord, pageSize int) int {
	return int(math.Ceil(float64(totalRecord) / float64(pageSize)))
}
