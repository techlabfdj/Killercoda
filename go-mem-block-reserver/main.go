package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	segmentRepo := NewSegmentRepository()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!\n")
	})

	r.GET("/gc", func(c *gin.Context) {
		runtime.GC()
		c.String(200, "GC Done !")
	})

	r.POST("/segments", func(c *gin.Context) {
		createSegment(c, segmentRepo)
	})

	r.GET("/segments/:id", func(c *gin.Context) {
		getSegment(c, segmentRepo)
	})

	r.PUT("/segments/:id", func(c *gin.Context) {
		updateSegment(c, segmentRepo)
	})

	r.DELETE("/segments/:id", func(c *gin.Context) {
		deleteSegment(c, segmentRepo)
	})

	r.GET("/segments", func(c *gin.Context) {
		listSegments(c, segmentRepo)
	})
	r.GET("/stats", func(c *gin.Context) {
		c.JSON(200, GetStats())
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
