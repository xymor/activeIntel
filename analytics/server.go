package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
  "github.com/twinj/uuid"
  "log"
  "os"
  "strings"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

var mongoURI = os.Getenv("MONGOURI")

var session *mgo.Session

func main() {
  log.Println("Starting")
  r := gin.Default()
  setupMongo()

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
    saveVisit(NewVisit(cookie.Value, c.Request))
    c.JSON(200, gin.H{"sucess": out})
  })
  r.Run(":8000")
}

type Visit struct {
        ID bson.ObjectId `bson:"_id,omitempty"`
        Timestamp time.Time 
        IP string
        Agent string
        Cookie string
        Referer string
        Query string
}

func NewVisit(cookieValue string, r *http.Request) *Visit {
    v := new(Visit)
    v.Timestamp = time.Now()
    v.Agent = r.UserAgent()
    v.Referer = r.Referer()
    v.IP = strings.Split(r.RemoteAddr, ":")[0] //remember X-Forwarded-For
    v.Cookie = cookieValue
    v.Query = r.URL.Query().Encode()
    return v
}

func saveVisit(visit *Visit) {
  log.Println("saving visit")
  log.Println(visit)
  c := session.DB("mongo").C("visits")
  err := c.Insert(visit)
  if err != nil {
          log.Fatal(err)
  }
}

func setupMongo() {
        var err error
        session, err = mgo.Dial(mongoURI)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
}