package handler

import (
	"fmt"
	"http_server/esutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createTodo add a new todo
func MySearch(c *gin.Context) {
	fmt.Println("-------in func mysearch()-------")
	c.String(http.StatusOK, "-------in func mysearch()-------")

	fmt.Println(c.FullPath(), c.ClientIP())
	fmt.Println(c.Query("sortedType"))
	fmt.Println(c.Query("pageNumber"))
	fmt.Println(c.Query("pageSize"))
	fmt.Println(c.Query("searchType"))

	sortedType := c.Query("sortedType")
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")
	searchType := c.Query("searchType")
	fmt.Printf("sortedType:%s, pageNumber:%s,pageSize:%s, searchType:%s ", sortedType, pageNumber, pageSize, searchType)

	//String sortedType, Integer pageNumber, Integer pageSize, DataType searchType
	//test_db()
	//mysqlxx.Test_mysql()
	esutil.Es_search(sortedType, pageNumber, pageSize, searchType)
}

// createTodo add a new todo
func MyPost(c *gin.Context) {
	fmt.Println("-------in func myPOST()-------")
	c.String(http.StatusOK, "-------in func myPOST()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	esutil.Test_es()
}

// createTodo add a new todo
func MyTest(c *gin.Context) {
	fmt.Println("-------in func mytest()-------")
	c.String(http.StatusOK, "-------in func mytest()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	esutil.Es_test()
}
