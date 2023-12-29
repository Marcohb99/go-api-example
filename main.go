package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	srv := gin.New()

	// ENDPOINTS

	// hc
	srv.GET("/health", healthHandler)

	// hello
	srv.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})

	// ListenAndServe returns error
	log.Fatal(srv.Run(httpAddr))
}

// 2nd param ignore because hc endpoints do not need request data
func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "everything is ok!")
}
