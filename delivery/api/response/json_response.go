package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, data interface{}) AppHttpResponse {
	httpStatusCode, repo := NewSuccessMessage(data)
	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       repo,
	}
}

func NewErrorJsonResponse(c *gin.Context, err error) AppHttpResponse {
	fmt.Println("===>", err)
	httpStatusCode, repo := NewSuccessMessage(err)
	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       repo,
	}
}

func NewGlobalJsonResponse(c *gin.Context, httpStatusCode int, response Response) AppHttpResponse {
	return &JsonResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       response,
	}
}
