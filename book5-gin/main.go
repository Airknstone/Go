//I recommend that you start off using go vet as a required part of your automated build process and golint as part of your code review process (since golint might have false positives and false negatives, you can’t require your team to fix every issue it reports). Once you are used to their recommendations, try out golangci-lint and tweak its settings until it works for your team.
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run()
}
