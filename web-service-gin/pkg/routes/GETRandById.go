package routes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func RandRedis(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		arrayLength, err := rdb.JSONArrLen(c, dataType+"_array", "$").Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		id := rand.Intn(int(arrayLength[0])) + 1
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
