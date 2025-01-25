package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EndpointDefinition struct {
	Method       string    `json:"method"`
	ResponseCode int       `json:"response_code"`
	ResponseBody string    `json:"response_body"`
	CreatedAt    time.Time `json:"created_at"`
}

type CustomEndpointHandler struct {
	endpoints         map[string]*EndpointDefinition
	rateLimits        map[string]int
	rateLimitDuration time.Duration
	mutex             sync.Mutex
}

func NewCustomEndpointHandler() *CustomEndpointHandler {
	return &CustomEndpointHandler{
		endpoints:         make(map[string]*EndpointDefinition),
		rateLimits:        make(map[string]int),
		rateLimitDuration: time.Hour,
		mutex:             sync.Mutex{},
	}
}

func (h *CustomEndpointHandler) Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Use /custom to create endpoints"})
}

func (h *CustomEndpointHandler) CreateEndpointHandler(c *gin.Context) {
	clientIP := c.ClientIP()

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if count, ok := h.rateLimits[clientIP]; ok && count >= 3 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many attempts. Please try again later"})
		return
	}

	var reqBody struct {
		Method       string `json:"method"`
		ResponseCode int    `json:"response_code"`
		ResponseBody string `json:"response_body"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hash := uuid.New().String()[:8]

	endpoint := &EndpointDefinition{
		Method:       reqBody.Method,
		ResponseCode: reqBody.ResponseCode,
		ResponseBody: reqBody.ResponseBody,
		CreatedAt:    time.Now(),
	}

	h.endpoints[hash] = endpoint

	go func() {
		time.Sleep(time.Hour)
		h.mutex.Lock()
		defer h.mutex.Unlock()
		delete(h.endpoints, hash)
	}()

	if _, ok := h.rateLimits[clientIP]; !ok {
		h.rateLimits[clientIP] = 1
	} else {
		h.rateLimits[clientIP]++
	}

	c.JSON(http.StatusCreated, gin.H{"hash": hash})
}

func (h *CustomEndpointHandler) HandleCustomEndpoint(c *gin.Context) {
	hash := c.Param("hash")
	endpoint, ok := h.endpoints[hash]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found or expired"})
		return
	}

	if c.Request.Method != endpoint.Method {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var result map[string]interface{}

	err := json.Unmarshal([]byte(endpoint.ResponseBody), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(endpoint.ResponseCode, result)
}

func (h *CustomEndpointHandler) EditCustomEndpointHandler(c *gin.Context) {
	hash := c.Param("hash")
	endpoint, ok := h.endpoints[hash]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
		return
	}

	var reqBody struct {
		Method       string `json:"method"`
		ResponseCode int    `json:"response_code"`
		ResponseBody string `json:"response_body"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpoint.Method = reqBody.Method
	endpoint.ResponseCode = reqBody.ResponseCode
	endpoint.ResponseBody = reqBody.ResponseBody
	// endpoint.CreatedAt = time.Now() // Reset expiration timer

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint updated"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//router := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	handler := NewCustomEndpointHandler()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"Msg": "Endpoint not defined", "Tip": "Check if the endpoint route is correct"})
	})

	r.GET("/custom", handler.Default)
	r.POST("/custom", handler.CreateEndpointHandler)
	r.Any("/custom/:hash", handler.HandleCustomEndpoint)
	r.PUT("/editcustom/:hash", handler.EditCustomEndpointHandler)

	r.Run(":8082")
}
