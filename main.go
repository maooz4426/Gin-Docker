package main

import (
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	//リクエストが成功した場合に、JSONを返す
	c.JSON(200, gin.H{
		"message": "Hello, Gin!",
	})
}

func main() {
	//*gin.Engineをrに入れる
	r := gin.Default()
	//エンドポイントの指定
	r.GET("/", helloWorld)
	//サーバーを起動
	r.Run()
}
