package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type fizzbuzzBody struct {
	Count int `json:"count" binding:"required"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/fizzbuzz", func(c *gin.Context) {
		var reqBody fizzbuzzBody

		// Bind the request body to the User struct
		if err := c.BindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			  	"message": "count parameter is required",
				"status": 500,
			})

			return
		}

		if reqBody.Count < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "count parameter is under 0",
				"status": 500,
			})

			return
		}

		var returnMsg = ""
		if reqBody.Count % 15 == 0 {
			returnMsg = "FizzBuzz"
		} else if reqBody.Count % 5 == 0 {
			returnMsg = "FizzBuzz"
		} else if reqBody.Count % 3 == 0 {
			returnMsg = "FizzBuzz"
		}

		c.JSON(200, gin.H {
			"status": 200,
			"message": returnMsg,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":4000")
}
