package handler

import (
	"fmt"
	"http_server/MySQLUtil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createTodo add a new todo
func MyGet(c *gin.Context) {
	fmt.Println("-------in func myget()-------")
	c.String(http.StatusOK, "-------in func myget()-------")
	MySQLUtil.Test_mysql()
}
