package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// struct for http post request body
type fizzbuzzBody struct {
	Count int `json:"count" binding:"required"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/fizzbuzz", func(c *gin.Context) {
		var reqBody fizzbuzzBody

		// Bind the request body to the fizzbuzzBody struct
		if err := c.BindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			  	"message": "count parameter is required",
				"status": 500,
			})

			return
		}

		// check if sent parameter value is under 0
		if reqBody.Count < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "count parameter is under 0",
				"status": 500,
			})

			return
		}

		// get return message
		var returnMsg = ""
		if reqBody.Count % 15 == 0 { // if multiple of 3 and 5
			returnMsg = "FizzBuzz"
		} else if reqBody.Count % 5 == 0 { // if multiple of 5
			returnMsg = "Buzz"
		} else if reqBody.Count % 3 == 0 { // if multiple of 3
			returnMsg = "Fizz"
		}

		// return message
		c.JSON(200, gin.H {
			"status": 200,
			"message": returnMsg,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":4000") // running on 4000 port
}
