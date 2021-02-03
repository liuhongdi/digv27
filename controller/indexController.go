package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv27/global"
	"time"
)

type IndexController struct{}
func NewIndexController() IndexController {
	return IndexController{}
}

// @Summary 商品详情
// @Description 显示商品的详情
// @Produce json
// @Param goodsid path int true "商品id"
// @Success 200 {object} swag.ResultContGoods 成功后返回值
// @Router /index/goodsone/{goodsid} [get]
func (g *IndexController) GoodsOne(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success("success");
	return
}
// @Summary 商品列表
// @Description 显示商品列表
// @Produce json
// @Param categoryId path int true "分类id"
// @Success 200 {object} swag.ResultContGoodsList 成功后返回值
// @Router /index/goodslist [get]
func (g *IndexController) GoodsList(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success("success");
	return
}