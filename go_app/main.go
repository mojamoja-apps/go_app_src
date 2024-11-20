package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		// データを取得
		data, err := getData()
		if err != nil {
			// エラーレスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// JSONでデータを返す
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	r.Run()
}
