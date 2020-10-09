package main

import (
	useDB "day10/book/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func bookListHandler(c *gin.Context) {
	booklist, err := useDB.QueryBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"cone": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": booklist,
	})
}

func bookDelHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"cone": 1,
			"msg":  err,
		})
		return
	}
	err = useDB.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"cone": 1,
			"msg":  err,
		})
		fmt.Println("删除失败")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8000")
}

func bookInsertHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

func bookNewhandler(c *gin.Context) {
	var form useDB.Book
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"cone": 1,
			"msg":  err,
		})
		return
	}
	fmt.Println(form.Title, form.Price)
	err := useDB.InsertBook(form.Title, form.Price)
	if err != nil {
		fmt.Println("添加失败")
		c.JSON(http.StatusOK, gin.H{
			"cone": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "new_book.html", gin.H{
		"code": 0,
		"msg":  "添加成功",
	})
}

func main() {
	// 初始化数据库
	err := useDB.Init("192.168.8.52", "gouser", "123456", "gotest")
	if err != nil {
		fmt.Println("数据库初始化失败:", err)
		return
	}
	// 初始化连接
	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("./templates/*")
	// 查询所有图书
	r.GET("/", bookListHandler)
	r.GET("/book/delete", bookDelHandler)
	r.GET("/book/new", bookInsertHandler)
	r.POST("/book/new", bookNewhandler)
	r.Run(":8000")
}
