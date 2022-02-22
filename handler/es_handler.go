package handler

import (
	"fmt"
	"http_server/esxx"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createTodo add a new todo
func MySearch(c *gin.Context) {
	fmt.Println("-------in func mysearch()-------")
	c.String(http.StatusOK, "-------in func mysearch()-------")
	//test_db()
	//mysqlxx.Test_mysql()
	esxx.Es_search()
}
