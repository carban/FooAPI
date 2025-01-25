package routes

import (
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"net/http"
	"os"
	"strconv"
	models "web-service-gin/models"

	"github.com/golang/freetype"

	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
)

func Redis(dataType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("GET-" + dataType)
		limitStr := c.Query("limit")
		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
				return
			}
			if limit < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be greater than zero"})
				return
			}
		}
		obj, err := rdb.JSONGet(c, dataType+"_array", "$.[0:"+limitStr+"]").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data from Redis"})
			return
		}
		// WARNING
		// Yes... There is no better way to do this
		// There is no way to change resData dynamically, you must to define the type...this is because the order of the attributtes in responses
		if dataType == "songs" {
			resData := []models.Song{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "users" {
			resData := []models.User{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "posts" {
			resData := []models.Post{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "comments" {
			resData := []models.Comment{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "products" {
			resData := []models.Product{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "todos" {
			resData := []models.Todo{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		} else if dataType == "movies" {
			resData := []models.Movie{}
			if err := json.Unmarshal([]byte(obj), &resData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": resData})
		}
	}
}

func GeoRedis(category string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("GET-" + category)
		limitStr := c.Query("limit")
		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
				return
			}
			if limit < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be greater than zero"})
				return
			}
		}
		obj, err := rdb.JSONGet(c, category+"_array", "$.features[0:"+limitStr+"]").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Error getting data from Redis"})
			return
		}
		resData := []*geojson.Feature{}
		if err := json.Unmarshal([]byte(obj), &resData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Unmarshal error"})
			return
		}
		fc := geojson.FeatureCollection{}
		fc.Features = resData
		c.JSON(http.StatusOK, gin.H{"data": fc})
	}
}

func Img() gin.HandlerFunc {
	return func(c *gin.Context) {
		addOne("GET-IMG")
		width, _ := strconv.Atoi(c.Param("width"))
		if width <= 0 || width > 2000 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Width must be greater than 0 and less than 1000"})
			return
		}
		height, _ := strconv.Atoi(c.Param("height"))
		if height <= 0 || height > 2000 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "Height must be greater than 0 and less than 1000"})
			return
		}
		bgColor := c.Param("bgColor")
		if len(bgColor) != 6 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "bg color must be a hex value like: ff0000"})
			return
		}
		fontColor := c.Param("fontColor")
		if len(fontColor) != 6 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "font color must be a hex value like: ff0000"})
			return
		}
		text := c.Param("text")
		if len(text) <= 0 || len(text) > 10 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "text length must be greater than 0 and less than 10"})
			return
		}
		textLenght := len(text)

		// Create a new RGBA image
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// Set the background color
		bg, _ := parseHexColor(bgColor)
		draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)

		// // Load a font (ensure you have a .ttf file for this)
		fontBytes, err := os.ReadFile("./etc/ARIAL.TTF") // Adjust path to your font file
		if err != nil {
			c.String(http.StatusInternalServerError, "Font file not found")
			return
		}

		fontParsed, err := freetype.ParseFont(fontBytes)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to parse font")
			return
		}
		fontSize := int(math.Floor(float64(height) * .1))
		// Draw the text on the image
		fc := freetype.NewContext()
		fc.SetDPI(72)
		fc.SetFont(fontParsed)
		fc.SetFontSize(float64(height) * .1) // Adjust the font size as needed
		fc.SetClip(img.Bounds())
		fc.SetDst(img)
		fc.SetSrc(&image.Uniform{colorFromHex(fontColor)})

		half := width/2 - textLenght*fontSize/4

		// Set the text position
		pt := freetype.Pt(half, int(fc.PointToFixed(float64(height)/2)>>6)) // Adjust positioning
		_, err = fc.DrawString(text, pt)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to draw text")
			return
		}

		// Serve the image as a PNG
		c.Header("Content-Type", "image/png")
		png.Encode(c.Writer, img)
	}
}

// Helper function to parse a hex color (e.g., "ff00ff")
func parseHexColor(s string) (color.Color, error) {
	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	g, _ := strconv.ParseUint(s[2:4], 16, 8)
	b, _ := strconv.ParseUint(s[4:6], 16, 8)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}, nil
}

// Helper function to convert hex to color
func colorFromHex(hex string) color.Color {
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
