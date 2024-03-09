/*
 * @Author: Vincent Yang
 * @Date: 2024-03-09 08:46:40
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2024-03-09 11:27:26
 * @FilePath: /ClaudeProxy/main.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2024 by Vincent, All Rights Reserved.
 */
package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func proxyRequest(c *gin.Context) {
	// Read the body from the original request
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}

	// Target URL
	targetURL := "https://api.anthropic.com" + c.Request.RequestURI

	// Create a new request
	proxyReq, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy request"})
		return
	}

	// Copy all request headers
	for header, values := range c.Request.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Proxy request failed"})
		return
	}
	defer resp.Body.Close()

	// Check if the response is compressed
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gzip reader"})
			return
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// Read the (decompressed) response body
	responseBody, err := io.ReadAll(reader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Forward the response body and status code back to the client
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}

func main() {
	var port string

	// Check for command line -p flag
	flag.StringVar(&port, "p", "", "Port to listen on")
	flag.Parse()

	// If -p was not used, look for PORT environment variable
	if port == "" {
		port = os.Getenv("PORT")
	}

	// Default to port 8080 if neither -p nor PORT environment variable is set
	if port == "" {
		port = "8080"
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	// Forward all /v1/* API requests
	r.Any("/v1/*action", proxyRequest)

	// Start server
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
