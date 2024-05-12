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
		api.GET("/albums", routes.GetAlbums)
		api.GET("/albums/:id", routes.GetAlbumByID)
		api.POST("/albums", routes.PostAlbums)
		api.GET("/geo", func(ctx *gin.Context) {
			fc := geojson.NewFeatureCollection()
			fc.AddFeature(geojson.NewPointFeature([]float64{1, 2}))
			fc.AddFeature(geojson.NewPointFeature([]float64{3, 4}))
			rawJSON, err := fc.MarshalJSON()
			fmt.Println(string(rawJSON))
			if err != nil {
				return
			}
			ctx.IndentedJSON(http.StatusOK, fc)
		})
		api.GET("/redis", routes.Redis("albums"))
		api.GET("/users", routes.Redis("users"))
		api.GET("/posts", routes.Redis("posts"))
		api.GET("/users/:id", routes.RedisById("users"))
	}

	router.Run("localhost:8080")
}
