package router

import (
	"datahub_exporter/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	route := gin.Default()
	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许的源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	// 使用 CORS 中间件
	route.Use(cors.New(config))

	routeApi := route.Group("/api")
	routeApiV1 := routeApi.Group("/v1")

	// 第三方
	routeApiV1CvmGroup := routeApiV1.Group("/ckafka")
	routeApiV1CvmGroup.GET("/connect_resources/metrics", controllers.CkafkaController{}.CkafkaDescribeConnectResourcesMetrcis)
	routeApiV1CvmGroup.GET("/datahub_tasks/metrics", controllers.CkafkaController{}.CkafkaDescribeDatahubTasksMetrcis)

	return route
}
