package swag

import "github.com/liuhongdi/digv27/model"

//商品详情
type ResultContGoods struct {
	Code int  `json:"code" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *model.Goods `json:"data" `  //返回的数据
}

//商品列表
type ResultContGoodsList struct {
	Code int  `json:"code" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data []*model.Goods `json:"data" `  //商品列表
}
