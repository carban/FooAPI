package main

import (
	"fmt"
	"net/http"
	"web-service-gin/pkg/routes"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//router := gin.New()
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")
	{
		api.GET("/albums", routes.GetAlbums)
		api.GET("/albums/:id", routes.GetAlbumByID)
		api.POST("/albums", routes.PostAlbums)
		api.GET("/users", routes.GetUsers)
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
	}

	router.Run("localhost:8080")
}
