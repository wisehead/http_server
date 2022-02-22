package main

import (
	"net/http"

	"http_server/handler"
	"http_server/mysqlxx"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/xxxpost", handler.MyGet)
	r.POST("/xxxpost", handler.MyPost)
	r.PUT("/xxxput", handler.MyGet)
	r.GET("/search", handler.MySearch)
	r.GET("/test", handler.MyTest)

	defer mysqlxx.Db.Close() // 注意这行代码要写在上面err判断的下面
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
