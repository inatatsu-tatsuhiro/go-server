package handler

import (
    // "net/http"

    // "api/article"

    "github.com/gin-gonic/gin"
)

func ArticleGet() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "article get",
        })
    }
}