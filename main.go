package main

import (
	"github.com/liuhongdi/digv27/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/liuhongdi/digv27/docs"
)

// @title 电商接口站文档
// @version 1.0
// @description 电商平台，api接口站的文档内容.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 371125307@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.estorelhd.com
func main() {
	//引入路由
	r := router.Router()
	//swagger:
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//run
	r.Run(":8080")
}
