# CORS Middleware for Gin-Gonic Framework

gin-middleware-cors is a middleware written in [Go (Golang)](http://golang.org) specifically for the [Gin Framework](https://gin-gonic.github.io/gin/) that implements the [Cross Origin Resource Sharing specification](http://www.w3.org/TR/cors/) from the W3C.  Implementing CORS headers enable pages within a modern web browser to consume resources (such as REST APIs) from servers that are on a different domain.

## Getting Started
To use this library, add the following code into your Gin router setup:

```go
import "github.com/benpate/go-cors"

// Initialize a new Gin router	
router := gin.New()

// Apply the middleware to the router (works with groups too)
router.Use(cors.Middleware(cors.Config{
	Origins:     "*",
	Headers:     "Origin, Authorization, Content-Type",
	Methods:     "GET, PUT, POST, DELETE",
	Credentials: true,
}))
```

## Setup Options
The middleware can be configured with four options, which match the HTTP headers that it generates:

Parameter   | Type   | Details
------------|--------|----------------------------------
Origins     | *string* | A comma+space separated list of allowable origins that is passed to the browser in the **Access-Control-Allow-Origins** header. For example: ```"http://localhost, http://api.server.com, http://files.server.com"```
Headers     | *string* | A comma+space separated list of allowable HTTP header codes that is passed to the browser in the **Access-Control-Allow-Headers** header.  Currently defaults to ```"Cache-Control, Pragma, Origin, Authorization, Content-Type"```
Methods     | *string* | A comma+space separated list of allowable HTTP methods that is passed to the browser in the **Access-Control-Allow-Methods** header.  Currently defaults to ```"GET, POST, PUT, DELETE"```
Credentials | *bool*   | A boolean value (true or false) that is passed in the **Access-Control-Allow-Credentials** header.  If ```true```, then cookies and Authorization headers are allowed in requests to this resource.  Defaults to ```false```.

## Caveats
I'm a new golang developer, and welcome your feedback about creating idiomatic Go middleware.  This code is highly experimental for me.  While I'm expecting to use it in large-scale web applications, it is almost certain to undergo massive changes to the API before that happens.

## CORS Resources

* [HTML Rocks Tutorial: Using CORS](http://www.html5rocks.com/en/tutorials/cors/)
* [Mozilla Developer Network: CORS Reference](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS)
* [CORS Specification from W3C](http://www.w3.org/TR/cors/)

## License
This code is licensed for free under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0). Copyright &copy; 2015 Ben & Fil LLC
