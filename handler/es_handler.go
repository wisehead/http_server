package handler

import (
	"fmt"
	"http_server/EsUtil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createTodo add a new todo
func MySearch(c *gin.Context) {
	fmt.Println("-------in func mysearch()-------")
	c.String(http.StatusOK, "-------in func mysearch()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	EsUtil.Es_search()
}

// createTodo add a new todo
func MyPost(c *gin.Context) {
	fmt.Println("-------in func myPOST()-------")
	c.String(http.StatusOK, "-------in func myPOST()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	EsUtil.Test_es()
}

// createTodo add a new todo
func MyTest(c *gin.Context) {
	fmt.Println("-------in func mytest()-------")
	c.String(http.StatusOK, "-------in func mytest()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	EsUtil.Es_test()
}
