package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deadlock-delegate/payfee-api/config"
	"github.com/deadlock-delegate/payfee-api/handlers"
	"github.com/deadlock-delegate/payfee-api/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
)

var allowOriging string

func main() {
	flag.StringVar(&config.Webhook.Token, "token", "", "Webhook token")
	flag.StringVar(&config.Webhook.VerificationKey, "verification", "", "Webhook verification key")
	flag.StringVar(&allowOriging, "allowOriging", "http://localhost:8080", "Access control allow origin")
	flag.Parse()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Println("Failed to connect to postgres:", err)
		os.Exit(1)
	} else {
		log.Println("Connected to postgres")
	}
	defer db.Close()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(func(c *gin.Context) {
		// todo: cleanup these values
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOriging)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")
		c.Next()
	})
	r.Use(gin.Recovery())

	socket, err := server.InitSocket()
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/ping", handlers.Ping)
	r.GET("/list", handlers.ListLinks)
	r.GET("/link/:linkID", handlers.GetLink)
	r.POST("/create", handlers.CreateLink)
	r.POST("/events", handlers.WebhookEvents)

	r.GET("/socket.io/*any", gin.WrapH(socket))
	r.POST("/socket.io/*any", gin.WrapH(socket))

	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
