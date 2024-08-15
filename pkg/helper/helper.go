package helper

import (
	"role-management/pkg/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RowData struct {
	Cells []string
}

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ConvertFloat(str string) float64 {
	valueStr := strings.ReplaceAll(str, ",", "")

	valueStr = strings.TrimSpace(valueStr)

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0.0
	}
	return value
}

// remove space ` f_name    l_name ` to `f_name l_name`
func TrimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func GetPagination(c *gin.Context) (*response.Pagination, error) {
	defaultPage := 1
	defaultPageSize := 10

	page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	if err != nil || page < 1 {
		page = defaultPage
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", strconv.Itoa(defaultPageSize)))
	if err != nil || pageSize < 1 {
		pageSize = defaultPageSize
	}

	return &response.Pagination{
		PageNumber: page,
		PageSize:   pageSize,
	}, nil
}
