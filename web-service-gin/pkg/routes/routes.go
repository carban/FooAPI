package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	models "web-service-gin/models"

	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

var rdb *redis.Client // Global Redis client

func init() {
	// Initialize Redis client in init function
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Album{})
}

func RedisAlbum(c *gin.Context) {
	obj, err := rdb.JSONGet(c, "albums_array").Result()
	if err != nil {
		c.IndentedJSON(http.StatusOK, "error")
		return
	}

	resData := []models.Album{}
	//xvar resData map[string]any
	err = json.Unmarshal([]byte(obj), &resData)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "error")
		return
	}
	c.IndentedJSON(http.StatusOK, resData)
}

func Redis(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {

		obj, err := rdb.JSONGet(c, dataType+"_array").Result()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data from Redis"})
			return
		}
		// WARNING
		// Yes... There is no better way to do this
		// There is no way to change resData dynamically, you must to define the type
		if dataType == "albums" {
			resData := []models.Album{}
			err = json.Unmarshal([]byte(obj), &resData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "users" {
			resData := []models.User{}
			err = json.Unmarshal([]byte(obj), &resData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "posts" {
			resData := []models.Post{}
			err = json.Unmarshal([]byte(obj), &resData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"data": resData})
		}
	}
}

func RedisById(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		obj, err := rdb.JSONGet(c, dataType+"_array", "["+fmt.Sprintf("%d", id-1)+"]").Result()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data", "Tip": "Check if the index is correct"})
			return
		}
		// WARNING
		// Yes... There is no better way to do this
		// There is no way to change resData dynamically, you must to define the type
		if dataType == "albums" {
			resData := models.Album{}
			err = json.Unmarshal([]byte(obj), &resData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "users" {
			resData := models.User{}
			err = json.Unmarshal([]byte(obj), &resData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"data": resData})
		}
	}
}

func RedisAlbumByID(c *gin.Context) {
	id := c.Param("id")
	obj, err := rdb.JSONGet(c, "albums_array", "["+id+"]").Result()
	if err != nil {
		c.IndentedJSON(http.StatusOK, "error")
		return
	}
	resData := models.Album{}
	//xvar resData map[string]any
	err = json.Unmarshal([]byte(obj), &resData)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "error")
		return
	}
	c.IndentedJSON(http.StatusOK, resData)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	// models.Albums = append(models.Albums, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	// for _, a := range models.Albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
