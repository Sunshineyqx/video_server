package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./video/")

	// 加载模板文件
	r.LoadHTMLFiles("templates/index.html", "templates/video.html")

	// 首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// qiqi
	r.GET("/qiqi", func(c *gin.Context) {
		videoPath := "qiqi.mp4"
		absolutePath, err := filepath.Abs(videoPath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %v", err)
			return
		}

		c.HTML(http.StatusOK, "video.html", gin.H{
			"videoPath": "/static/" + filepath.Base(absolutePath),
			"name":      "可爱的丸丸琪",
		})
	})

	// waxing
	r.GET("/waxing", func(c *gin.Context) {
		videoPath := "waxing.mp4" // 假设视频文件名为 video.mp4，放在主目录下
		absolutePath, err := filepath.Abs(videoPath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %v", err)
			return
		}

		c.HTML(http.StatusOK, "video.html", gin.H{
			"videoPath": "/static/" + filepath.Base(absolutePath),
			"name":      "春春de娃行",
		})
	})
	r.Run(":8080")
}
