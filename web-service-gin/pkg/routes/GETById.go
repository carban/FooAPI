package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func RedisById(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, convErr := strconv.Atoi(c.Param("id"))
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be a number"})
			return
		}
		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be bigger than 0"})
			return
		}
		obj, err := rdb.JSONGet(c, dataType+"_array", "["+fmt.Sprintf("%d", id-1)+"]").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		// WARNING
		// Yes... There is no better way to do this
		// There is no way to change resData dynamically, you must to define the type...this is because the order of the attributtes in responses
		if dataType == "songs" {
			resData := models.Song{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "users" {
			resData := models.User{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "posts" {
			resData := models.Post{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "comments" {
			resData := models.Comment{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "products" {
			resData := models.Product{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "todos" {
			resData := models.Todo{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "movies" {
			resData := models.Movie{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		}
	}
}

func GeoRedisById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, convErr := strconv.Atoi(c.Param("id"))
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be a number"})
			return
		}
		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be bigger than 0"})
			return
		}
		obj, err := rdb.JSONGet(c, "cities_array", ".features["+fmt.Sprintf("%d", id-1)+"]").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		resData := geojson.Feature{}
		if err := json.Unmarshal([]byte(obj), &resData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": resData})
	}
}
