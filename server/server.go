package main

import (
	"fmt"
	"net/http"

	"http_server/mysqlxx"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close() // 注意这行代码要写在上面err判断的下面
}

func test_db() {
	if Db == nil {
		fmt.Println("Db is nil")
	}
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

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
	mysqlxx.Test_mysql()
}

// createTodo add a new todo
func mypost(c *gin.Context) {
	fmt.Println("-------in func myPOST()-------")
	c.String(http.StatusOK, "-------in func myPOST()-------")
	test_db()
}
