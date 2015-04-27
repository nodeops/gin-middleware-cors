package cors

import (
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() {

	// Initialize the gin-gonic router
	router := gin.New()

	// Set up CORS middleware options
	config := Config{
		Origins:     "*",
		Methods:     "GET, POST, PUT",
		Credentials: true,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(Middleware(config))
}
