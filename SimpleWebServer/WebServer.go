package main

import (
	"fmt"
	"net/http"
)

// HTML content as a string
const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple Web Server</title>
</head>
<body>
    <h1>Welcome to the Simple Web Server</h1>
    <p>This is a static HTML page served by a Go web server.</p>
</body>
</html>
`

func main() {
	// Handle the root URL by serving the HTML content
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)
	})

	// Define the port to listen on
	port := "8080"
	fmt.Printf("Starting server at port %s\n", port)

	// Start the server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
