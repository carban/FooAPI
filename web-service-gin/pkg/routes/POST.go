package routes

import (
	"net/http"
	"strconv"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func PostRedis(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("POST-" + dataType)
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
		// WARNING
		// Yes... There is no better way to do this
		// There is no way to change resData dynamically, you must to define the type...this is because the order of the attributtes in responses
		if dataType == "songs" {
			var newSong models.SongReq
			if err := c.BindJSON(&newSong); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if !newSong.IsExplicit {
				newSong.IsExplicit = false
			}
			newSong.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newSong})
		} else if dataType == "users" {
			var newUser models.UserReq
			if err := c.BindJSON(&newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newUser.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newUser})
		} else if dataType == "posts" {
			var newPost models.PostReq
			if err := c.BindJSON(&newPost); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newPost.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newPost})
		} else if dataType == "comments" {
			var newComment models.CommentReq
			if err := c.BindJSON(&newComment); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newComment.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newComment})
		} else if dataType == "products" {
			var newProduct models.ProductReq
			if err := c.BindJSON(&newProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newProduct.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newProduct})
		} else if dataType == "todos" {
			var newTodo models.TodoReq
			if err := c.BindJSON(&newTodo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if !newTodo.Closed {
				newTodo.Closed = false
			}
			newTodo.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newTodo})
		} else if dataType == "movies" {
			var newMovie models.MovieReq
			if err := c.BindJSON(&newMovie); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newMovie.ID = strconv.Itoa(len + 1)
			c.JSON(http.StatusCreated, gin.H{"data": newMovie})
		}
	}
}

func GeoRedisPost(category string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("POST-" + category)
		var newFeature geojson.Feature
		if err := c.BindJSON(&newFeature); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": newFeature})
	}
}
