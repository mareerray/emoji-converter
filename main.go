package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)


func main() {

	// Handle the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var emoji string
		var inputText string

			// Template for the error page
		errorTmpl := template.Must(template.ParseFiles("error.html"))
		error500Tmpl := template.Must(template.ParseFiles("error500.html"))

		// Template for the HTML page
		// tmpl:= template.Must(template.ParseFiles("emojiConverter.html"))
		tmpl, err := template.ParseFiles("emojiConverter.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			error500Tmpl.Execute(w, struct {
				Status int
				Message string
			}{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
			return
			// fmt.Printf("Error loading main template: %v\n", err)
			// return
		}

		// Handle invalid paths with a custom 404 error page
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			errorTmpl.Execute(w, struct {
				Status int
				Message string
			}{
				Status:  http.StatusNotFound,
                Message: "Page Not Found",
			})
			return
		}		
	
		// Handle GET requests
		if r.Method == http.MethodGet {
			// Render the template with empty input and output
			tmpl.Execute(w, struct {
				InputText string
				Emoji     string
			}{
				InputText: "",
				Emoji:     "",
			})
			return
		}

		// Handle POST requests
		if r.Method == http.MethodPost {
			r.ParseForm()
			inputText = strings.ToLower(r.PostForm.Get("input-text"))

			// Check if the input text is empty
			if inputText == "" {
				emoji = "⚠️Please enter a valid emoji name"
			} else {
				// Check if the input text is in the emoji map
				// If it is, get the corresponding emoji
				emoji = EmojiMap[inputText]
				// If the emoji is not found, set a default message
				if emoji == "" {
					emoji = "❓Emoji not found"
				}
			}
		}
		// Execute the template and pass the emoji to it
		tmpl.Execute(w, struct {
			InputText string 
			Emoji string
		}{
			InputText: inputText,
			Emoji: emoji})
	})
	// Start the HTTP server on port 8080
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Default behavior: Handle GET (or other methods) by rendering the form
	// No need to block non-POST requests
	// Allow GET requests to load the page initially.
	// Process POST requests only when the form is submitted.
	// The template renders whether the request is GET or POST.

// HTML Flow

// sequenceDiagram
//     Browser->>Server: GET / (Initial Load)
//     Server->>Browser: Render empty form
//     Browser->>Server: POST / (Submit form)
//     Server->>Browser: Render form with results

// Testing
// 1) Initial Page Load (GET):

// - Visit http://localhost:8080.

// - You’ll see the form with empty input and output fields.

// 2) Form Submission (POST):

// - Enter "rocket" and click "Convert".

// - The input field retains "rocket", and the output shows "🚀".

// 3) Edge Case:

// - Enter an unknown name like "unknown".

// - Output shows "❓Emoji not found".

// Why This Works
// - GET Requests: Serve the initial form.

// - POST Requests: Process input and update results.

// - The same template handles both scenarios.

// This approach ensures your application works seamlessly for both initial page loads and form submissions! 