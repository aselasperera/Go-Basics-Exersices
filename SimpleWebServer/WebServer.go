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
    <h1>Welcome to the Evolza Web Server</h1>
    <p>This is a static HTML page served by a Go web server.</p>
</body>
</html>
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)
	})

	port := "8080"
	fmt.Printf("Starting server at port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
