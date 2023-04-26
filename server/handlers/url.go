package handlers

import (
	"fmt"
)

var (
	server string
	port   string
)

// SetConfig sets the server and port to use for HTTP requests
func SetConfig(s string, p string) {
	server = s
	port = p
}

// GenerateURL generates a URL using the provided endpoint and configuration
func GenerateURL(endpoint string) string {
	return fmt.Sprintf("http://%s:%s/%s", server, port, endpoint)
}
