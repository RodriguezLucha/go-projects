package main

import "github.com/gin-gonic/gin"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func get_ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pongg",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", get_ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
