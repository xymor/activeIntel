package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
  "github.com/twinj/uuid"
)

func main() {
  r := gin.Default()

  r.GET("/track", func(c *gin.Context) {
  	
	cookie, err := c.Request.Cookie("ecommtracker")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "ecommtracker",
			Expires: time.Now().Add(10 * 365 * 24 * time.Hour),
			Value: uuid.NewV4().String(),
			Path:  "/",
		} 
  		http.SetCookie(c.Writer, cookie)
	}
	out := cookie.Value		

    c.JSON(200, gin.H{"sucess": out})
  })
  r.Run(":8000")
}