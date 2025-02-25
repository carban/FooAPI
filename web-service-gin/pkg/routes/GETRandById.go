package routes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func RandRedis(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("GETRand-" + dataType)
		arrayLength, err := rdb.Get(c, dataType+"_len").Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		len, convErr := strconv.Atoi(arrayLength)
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing the data"})
			return
		}
		id := rand.Intn(len) + 1
		obj, err := rdb.JSONGet(c, dataType+"_array", "["+fmt.Sprintf("%d", id-1)+"]").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
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

func GeoRedisRand(category string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("GETRand-" + category)
		arrayLength, err := rdb.Get(c, category+"_len").Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		len, convErr := strconv.Atoi(arrayLength)
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing the data"})
			return
		}
		id := rand.Intn(len) + 1
		obj, err := rdb.JSONGet(c, category+"_array", ".features["+fmt.Sprintf("%d", id-1)+"]").Result()
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
