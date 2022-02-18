package main

import (
	"fmt"
	"net/http"

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
	r.GET("/xxxpost", myget)
	r.POST("/xxxpost", mypost)
	r.PUT("/xxxput", myget)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

// createTodo add a new todo
func myget(c *gin.Context) {
	fmt.Println("-------in func myget()-------")
	c.String(http.StatusOK, "-------in func myget()-------")
}

// createTodo add a new todo
func mypost(c *gin.Context) {
	fmt.Println("-------in func myPOST()-------")
	c.String(http.StatusOK, "-------in func myPOST()-------")
}
