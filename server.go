package main

import (
    "github.com/gin-gonic/gin"
    // "api/article"
    "api/handler"
)


func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/article", handler.ArticleGet())
    r.Run(":8080")
}