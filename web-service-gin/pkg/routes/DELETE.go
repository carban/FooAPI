package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRedisById(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, convErr := strconv.Atoi(c.Param("id"))
		if convErr != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be a number"})
			return
		}
		if id <= 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Msg": "Error parsing the data", "Tip": "The Id must be bigger than 0"})
			return
		}
		obj, err := rdb.JSONGet(c, dataType+"_array", "["+fmt.Sprintf("%d", id-1)+"].id").Result()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		var resData string
		err = json.Unmarshal([]byte(obj), &resData)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"data": gin.H{"id": resData}})
	}
}
