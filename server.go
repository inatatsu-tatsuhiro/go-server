package main

import (
    "github.com/gin-gonic/gin"
    "api/article"
    "api/handler"
    "fmt"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)


func main() {
    db := sqlConnect()
    defar db.close()
    article := article.New()
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        _, err := sqlConnect()
        if err != nil {
            c.JSON(200, gin.H{
                "message": err,
            })
        } else {
            c.JSON(200, gin.H{
                "message": "ping",
            })
        }
    })
    r.GET("/article", handler.ArticlesGet(article))
    r.POST("/article", handler.ArticlePost(article))

    r.Run(":8080")
}

func index(w http.ResponseWriter, r *http.Request) {
    _, err := sqlConnect()
    if err != nil {
        fmt.Fprintf(w, err.Error())
    } else {
        fmt.Fprintf(w, "DB接続成功")
    }
}

func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "docker"
    PROTOCOL := "tcp(db:3306)"
    DBNAME := "docker"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}