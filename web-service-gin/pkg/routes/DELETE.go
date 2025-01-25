package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRedisById(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("DELETE-" + dataType)
		id, convErr := strconv.Atoi(c.Param("id"))
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be a number"})
			return
		}
		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be bigger than 0"})
			return
		}
		obj, err := rdb.JSONGet(c, dataType+"_array", "["+fmt.Sprintf("%d", id-1)+"].id").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		var resData string
		err = json.Unmarshal([]byte(obj), &resData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": resData}})
	}
}

func GeoRedisDeleteById(category string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("DELETE-" + category)
		var (
			sid    string
			err    error
			errMsg string
		)

		if id, convErr := strconv.Atoi(c.Param("id")); convErr == nil {
			// by index
			if id <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be bigger than 0"})
				return
			}
			sid = ".features[" + fmt.Sprintf("%d", id-1) + "].id"
		} else {
			// by ID
			if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(c.Param("id")) {
				c.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid ID", "Tip": "Check the ID"})
				return
			}
			sid = fmt.Sprintf(".features[?(@.id==\"%s\")].id", c.Param("id"))
		}
		obj, err := rdb.JSONGet(c, category+"_array", sid).Result()
		if err != nil {
			if _, err := strconv.Atoi(c.Param("id")); err == nil {
				errMsg = "Check if the index is correct"
			} else {
				errMsg = "Check if the index/ID is correct. Remember ID must be use capital letters"
			}
			c.JSON(http.StatusBadRequest, gin.H{"Msg": "Error getting data", "Tip": errMsg})
			return
		}
		var resData string
		if err := json.Unmarshal([]byte(obj), &resData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": resData})
	}
}
