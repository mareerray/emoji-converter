# ✨Emoji Converter✨	

Emoji Converter Description
The Emoji Converter is a simple and interactive webpage that allows users to convert emoji names into their corresponding symbols. Enter an emoji name (e.g., "smile," "heart," "rocket") into the input field, click "Convert," and see the magic happen! If the name is invalid or empty, helpful error messages guide you. Perfect for quick emoji lookups and fun exploration of emojis! 


## How to Run the Application

1. Clone the repository:
    ```
    git clone https://github.com/your-username/emoji-converter.git
    cd emoji-converter
    ```

2. Install dependencies:
    ```
    go mod tidy
    ```

3. Run the application:
    ```
    go run main.go
    ```

4. Open your browser and navigate to:
    ```
    http://localhost:8080/
    ```

## Instructions for Using the Emoji Converter Webpage
1) Access the Webpage:

- Open your browser and go to http://localhost:8080 after running the application.

2) Enter an Emoji Name:

- In the input box, type the name of an emoji (e.g., smile, heart, rocket, etc.).

3) Click "Convert" or Enter:

- Press the "Convert" button or Enter to see the corresponding emoji appear in the output container.

4) If Input is Empty:

- If you leave the input field blank and click "Convert," you will see a message prompting you to enter a valid emoji name.

5) If Emoji is Not Found:

- If the name you entered does not match any available emoji, a message saying "❓Emoji not found" will appear.

6) Navigate Back to Homepage:

- If you encounter a "Page Not Found" error (404), click the "Go back to Homepage" button to return to the main page.

7) Error Handling:

- If there’s an unexpected issue, "500 - Internal Server Error" will be displayed.

## Features

- Converts emoji names (e.g., "smile") into their corresponding emoji symbols (e.g., 😀).
- Custom 404 and 500 error pages.
- Responsive design using inline CSS.


## Future Improvements

- Add support for more emojis.
- Implement a database to store user input history.
- Add unit tests for core functionality.

## project structure
```
Emoji Converter project
├── index.html          # Main HTML file
├── main.go             # Entry point for the Go application
├── error.html          # Custom 404 error page
├── error500.html       # Custom 500 error page
├── emojis.go           # Emoji mapping 
├── go.mod              # Go module file
├── README.md           # Project documentation

```

## Creator: 
Mayuree Reunsati (April / 2025)