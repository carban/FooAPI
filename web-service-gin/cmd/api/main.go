package main

import (
	"net/http"
	"os"
	"strings"
	"web-service-gin/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	//router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{"Msg": "Endpoint not defined", "Tip": "Check if the endpoint route is correct"})
		}
	})

	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"Msg": "Welcome to FooApi"})
		})

		categories := []string{"songs", "users", "posts", "comments", "products", "todos", "movies"}

		for _, category := range categories {
			api.GET("/"+category, routes.Redis(category))
			api.GET("/"+category+"/:id", routes.RedisById(category))
			api.GET("/"+category+"/rand", routes.RandRedis(category))
			api.PUT("/"+category+"/:id", routes.PutRedisById(category))
			api.PATCH("/"+category+"/:id", routes.PatchRedisById(category))
			api.POST("/"+category, routes.PostRedis(category))
			api.DELETE("/"+category+"/:id", routes.DeleteRedisById(category))
		}

		geoCategories := []string{"cities", "countries"}

		for _, category := range geoCategories {
			api.GET("/"+category, routes.GeoRedis(category))
			api.GET("/"+category+"/:id", routes.GeoRedisById(category))
			api.GET("/"+category+"/rand", routes.GeoRedisRand(category))
			api.PUT("/"+category+"/:id", routes.GeoRedisPutById(category))
			api.PATCH("/"+category+"/:id", routes.GeoRedisPatchById(category))
			api.POST("/"+category, routes.GeoRedisPost(category))
			api.DELETE("/"+category+"/:id", routes.GeoRedisDeleteById(category))
		}
	}

	router.GET("/imgmaker/:width/:height/:bgColor/:fontColor/:text", routes.Img())
	router.GET("/img/:id", func(c *gin.Context) {
		id := c.Param("id")
		filePath := "./images/" + id

		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		c.File(filePath)
	})

	router.Run("localhost:8080")
}
