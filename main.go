package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	secret := "m9kaxFePwArUwRs53qaOsSoFP6bjpFD6"
	hub := newHub()
	go hub.run()

	router := gin.New()
	router.LoadHTMLGlob("public/*.html")

	router.GET("/room/:roomID", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		ipAdds := c.ClientIP()
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(ipAdds))
		userToken := hex.EncodeToString(h.Sum(nil))

		serveWs(hub, c.Writer, c.Request, roomID, userToken)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
