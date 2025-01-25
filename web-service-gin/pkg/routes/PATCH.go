package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func PatchRedisById(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("PATCH-" + dataType)
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
			var newSong models.Song
			if err := c.ShouldBind(&newSong); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Song{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newSong).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			if !newSong.IsExplicit {
				resData.IsExplicit = false
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "users" {
			var newUser models.User
			if err := c.ShouldBind(&newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.User{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newUser).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "posts" {
			var newPost models.Post
			if err := c.ShouldBind(&newPost); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Post{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newPost).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "comments" {
			var newComment models.Comment
			if err := c.ShouldBind(&newComment); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Comment{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newComment).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "products" {
			var newProduct models.Product
			if err := c.ShouldBind(&newProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Product{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newProduct).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "todos" {
			var newTodo models.Todo
			if err := c.ShouldBind(&newTodo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Todo{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newTodo).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			if !newTodo.Closed {
				resData.Closed = false
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		} else if dataType == "movies" {
			var newMovie models.Movie
			if err := c.ShouldBind(&newMovie); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resData := models.Movie{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			v := reflect.ValueOf(&resData).Elem()
			t := v.Type()

			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldName := field.Name
				fieldValue := reflect.ValueOf(newMovie).FieldByName(fieldName)

				if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
					v.Field(i).Set(fieldValue)
				}
			}
			c.JSON(http.StatusCreated, gin.H{"data": resData})
		}
	}
}

func GeoRedisPatchById(category string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("PATCH-" + category)
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
			sid = ".features[" + fmt.Sprintf("%d", id-1) + "]"
		} else {
			// by ID
			if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(c.Param("id")) {
				c.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid ID", "Tip": "Check the ID"})
				return
			}
			sid = fmt.Sprintf(".features[?(@.id==\"%s\")]", c.Param("id"))
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
		var newFeature geojson.Feature
		if err := c.ShouldBind(&newFeature); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resData := geojson.Feature{}
		if err := json.Unmarshal([]byte(obj), &resData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		v := reflect.ValueOf(&resData).Elem()
		t := v.Type()

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Name
			fieldValue := reflect.ValueOf(newFeature).FieldByName(fieldName)

			if fieldValue.IsValid() && !fieldValue.IsZero() && fieldName != "ID" {
				v.Field(i).Set(fieldValue)
			}
		}
		c.JSON(http.StatusCreated, gin.H{"data": resData})

	}
}
