package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

type ResultCont struct {
	Code int  `json:"code" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data interface{} `json:"data" `  //返回的数据
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

//返回成功
func (r *Result) Success(data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK,res)
}

//返回失败
func (r *Result)Error(code int,msg string) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()
}
