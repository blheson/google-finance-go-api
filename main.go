package main

import (
	"net/http"
	"os"

	"github.com/blheson/project/request"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Index
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, request.Greet("Welcome to BlessingUdor Finance GO Api"))
	})

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })
	r.GET("/search", func(c *gin.Context) {

		// log.Print("==========>", "/search?q="+c.Query("q"), "========>")
		//TODO: handle other keys in query
		body, err := request.Get("/search/?q=" + c.Query("q"))

		if err != nil {
			print(err)

			return
		}
		// gin.H{"body": string(body)}
		// fmt.Println("\n" + )
		// fmt.Println(string(body))
		c.JSON(http.StatusOK, string(body))

	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
func init() {
	gotenv.Load()
}
func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PATH")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
