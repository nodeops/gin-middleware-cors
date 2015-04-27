// Copyright 2015  Ben & Fil LLC. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

/*
gin-middleware-cors is a middleware written specifically for the Gin Framework that implements the
Cross Origin Resource Sharing specification from the W3C.  Implementing CORS headers enable pages
within a modern web browser to consume resources (such as REST APIs) from servers that are on a
different domain.
*/
package cors

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
Config defines the configuration options available in the Middleware function.
*/
type Config struct {

	// Comma+space separated list of origin domains. Wildcard "*" is also allowed, and matches
	// all origins.  If the origin does not match an item in the list, then the request is denied.
	Origins string

	// Comma+space separated list of acceptable HTTP headers.  This is passed to the browser, but
	// is not enforced.  Default is "Cache-Control, Pragma, Origin, Authorization, Content-Type"
	Headers string

	// Comma+space separated list of acceptable HTTP methods.  This is passed to the browser, but
	// is not enforced.  Default is "GET, POST, PUT, DELETE"
	Methods string

	// If true, then cookies and Authorization headers are allowed along with the request.  This
	// is passed to the browser, but is not enforced.
	Credentials bool
}

/*
Middleware generates a middleware handler function that works inside of a Gin request
to set the correct CORS headers.  It accepts a cors.Options struct for configuration.
*/
func Middleware(config Config) gin.HandlerFunc {

	// Convert the list into a slice of strings.  This happens just once for the life of the server
	validOrigins := strings.Split(config.Origins, ", ")

	// Set default value for headers
	if config.Headers == "" {
		config.Headers = "Cache-Control, Pragma, Origin, Authorization, Content-Type"
	}

	// Set default value for methods
	if config.Methods == "" {
		config.Methods = "GET, POST, PUT, DELETE"
	}

	// Format boolean as a string (True/False)
	credentials := fmt.Sprintf("%t", config.Credentials)

	// Create the Middleware function
	return func(context *gin.Context) {

		// Read the Origin header from the HTTP request
		currentOrigin := context.Request.Header.Get("Origin")

		// CORS headers are added whenever the browser request includes an "Origin" header
		if currentOrigin != "" {

			// If the first item in the array is "*" then allow requests from any location.
			if validOrigins[0] == "*" {
				validOrigins = []string{currentOrigin}
			}

			// Loop over all items in the validOrigins
			for _, value := range validOrigins {

				// if we have a match, then add the required request headers and continue.
				if value == currentOrigin {

					context.Writer.Header().Set("Access-Control-Allow-Origin", currentOrigin)
					context.Writer.Header().Set("Access-Control-Allow-Headers", config.Headers)
					context.Writer.Header().Set("Access-Control-Allow-Methods", config.Methods)
					context.Writer.Header().Set("Access-Control-Allow-Credentials", credentials)

					// Exit here because we're successful
					return

				}
			}

			// Fall through to here means that the origin does not match.
			// Return an HTTP 403 (Forbidden) to tell the browser it is not welcome here.
			context.AbortWithStatus(403)
		}
	}
}
