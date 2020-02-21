package app


import (
	"bfs_cli_rsa/common/e"
	"github.com/gin-gonic/gin"

)



type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


// Response setting gin.JSON
func NewResponse(C *gin.Context,httpCode, errCode int, data interface{}) {
	C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func NewOriginResponse(errCode int, data interface{}) Response{
	return Response{Code:errCode,Msg:e.GetMsg(errCode),Data:data}
}
func NewDecorateResponse(C *gin.Context,httpStatusCode int, response Response) {
	C.JSON(httpStatusCode, Response{
		Code: response.Code,
		Msg:  e.GetMsg(response.Code),
		Data: response.Data,
	})
	return
}
