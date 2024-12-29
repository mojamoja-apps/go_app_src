package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go_app/model"
)

func main() {

	r := gin.Default()

	// CORS設定を追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/user", func(c *gin.Context) {
		// データを取得
		users, err := model.GetUsers()
		if err != nil {
			// エラーレスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// JSONでデータを返す
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})

	r.GET("/station", func(c *gin.Context) {
		// クエリパラメータを取得
		neLat, _ := strconv.ParseFloat(c.Query("ne_lat"), 64)
		neLng, _ := strconv.ParseFloat(c.Query("ne_lng"), 64)
		swLat, _ := strconv.ParseFloat(c.Query("sw_lat"), 64)
		swLng, _ := strconv.ParseFloat(c.Query("sw_lng"), 64)

		// データを取得
		stations, err := model.GetStations(neLat, neLng, swLat, swLng)
		if err != nil {
			// エラーレスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// JSONでデータを返す
		c.JSON(http.StatusOK, gin.H{
			"data": stations,
		})
	})
	r.Run()
}
