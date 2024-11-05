package main

import (
	"fmt"
	"net/http"
	"strings"
	"web-service-gin/pkg/routes"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
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

	api := router.Group("/api")
	{
		api.GET("/geo", func(ctx *gin.Context) {
			fc := geojson.NewFeatureCollection()
			fc.AddFeature(geojson.NewPointFeature([]float64{1, 2}))
			fc.AddFeature(geojson.NewPointFeature([]float64{3, 4}))
			rawJSON, err := fc.MarshalJSON()
			fmt.Println(string(rawJSON))
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, fc)
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
		api.GET("/capitals", routes.GeoRedis())
		api.GET("/capitals/:id", routes.GeoRedisById())
	}
	router.Run("localhost:8080")
}
