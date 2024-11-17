package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/version/guestbook/:gbId", GetGuestbook)
	router.POST("/api/version/guestbook", PostGuestbook)
	router.PUT("/api/version/guestbook", PutGuestbook)
	router.DELETE("/api/version/guestbook/:gbId", DeleteGuestbook)

	router.GET("/api/version/guestbook/:gbId/messages", GetMessagesByGuestbookId)
	router.POST("/api/version/guestbook/:gbId/message", PostMessage)
	router.PUT("/api/version/guestbook/:gbId/message", PutMessage)
	router.DELETE("/api/version/guestbook/:gbId/message/:msgId", DeleteMessage)

	router.Run(":8080")
}
