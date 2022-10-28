package Controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zxc10110/mvc_09/View"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/createStudent", CreateStudent)
	r.GET("/getStudent", GetStudent)
	r.GET("/randomSport", SportStart)
	r.GET("/getSport", GetSport)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
