package main

import "github.com/gin-gonic/gin"

func pingPong() {
	r := gin.Default()
	// 	gin.Default() creates a Gin router with
	// default middleware (such as logging and recovery from panics).
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 	r.GET("/ping", func(c *gin.Context) sets up a route
	// that listens for GET requests at /ping and responds with JSON.
	r.Run() // listen and serve on 0.0.0.0:8080
	/* router.Run(":3000") for a hard coded port*/
}

/*
The second argument, func(c *gin.Context),
is a callback function that gets called whenever a
GET request is made to /ping
The gin.Context object (c) is passed to the callback,
which contains all the information about the request
(headers, query parameters, etc.) and also provides methods to
write a response back to the client.
*/
/*In your code, *gin.Context is a pointer to a gin.Context struct.
c is a pointer to an instance of the gin.Context struct.Instead of passing the entire gin.Context struct (which could be a large object), only the memory address where the struct is stored is passed.
This is more efficient in terms of performance.*/
