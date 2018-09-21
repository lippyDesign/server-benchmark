package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func createPersons() []*Person {
	var persons []*Person
	for i := 0; i < 101; i++ {
		person := &Person{
			FirstName: randStringBytesMaskImpr(7),
			LastName:  randStringBytesMaskImpr(7),
			ID:        randStringBytesMaskImpr(15),
			Phone:     rand.Intn(10000000000),
			Street:    randStringBytesMaskImpr(15),
			City:      randStringBytesMaskImpr(10),
			State:     randStringBytesMaskImpr(2),
			Zip:       rand.Intn(100000),
		}
		persons = append(persons, person)
	}
	return persons
}

func main() {
	engine := gin.Default()

	// GET 'Hello World!' will send back 'hello world' as a string
	engine.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// GET static files
	engine.LoadHTMLGlob("./client/build/*.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// GET json blob- array of 100 random persons
	engine.GET("/persons", func(c *gin.Context) {
		Persons := createPersons()
		c.JSON(http.StatusOK, Persons)
	})

	// set up port the server will listen on
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	port = ":" + port
	engine.Run(port)
}
