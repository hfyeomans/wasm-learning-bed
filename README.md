# Go WebAssembly Number Adder

A simple WebAssembly project that demonstrates how to use Go code compiled to WebAssembly to perform calculations in the browser.

## Overview

This project showcases:
- Compiling Go code to WebAssembly (WASM)
- Running Go functions in the browser
- Two-way communication between JavaScript and Go
- Serving WebAssembly files with the correct MIME types

The application provides a simple web interface where users can enter two numbers, and the addition is performed by Go code running as WebAssembly in the browser.

## Project Structure

- `main.go` - The Go code that gets compiled to WebAssembly
- `server.go` - A simple HTTP server to serve the application
- `index.html` - The web interface with inputs for numbers
- `wasm_exec.js` - JavaScript glue code provided by Go to facilitate WebAssembly integration
- `main.wasm` - The compiled WebAssembly binary (generated from main.go)

## Requirements

- Go 1.24 or later (tested with Go 1.24.2)
- A modern web browser that supports WebAssembly

## Setup and Running

1. **Build the WebAssembly binary:**

   ```bash
   GOOS=js GOARCH=wasm go build -o main.wasm main.go
   ```

2. **Copy the WebAssembly support JavaScript file:**

   ```bash
   cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
   ```
   
   Note: In Go 1.24, the file is located at `$(go env GOROOT)/lib/wasm/wasm_exec.js`

3. **Build the server:**

   ```bash
   go build server.go
   ```

4. **Run the server:**

   ```bash
   ./server
   ```

5. **Open the application in your browser:**

   Navigate to [http://localhost:8080](http://localhost:8080)

## How It Works

### Go Side (main.go)

The Go code exposes an `Add` function to JavaScript using the `syscall/js` package. The function:
1. Receives two integers as arguments from JavaScript
2. Adds them together
3. Returns the result to JavaScript

### JavaScript Side (index.html)

When the page loads:
1. The WebAssembly module is fetched and instantiated
2. The Go runtime is initialized
3. The Go function `addNumbers` is registered and available to JavaScript

When the user clicks the "Add" button:
1. Input values are read from the form
2. The Go function is called with these values
3. The result is displayed on the page

### Server Side (server.go)

The server:
1. Serves static files from the current directory
2. Sets the correct MIME type for WebAssembly files (`application/wasm`)
3. Listens on port 8080 by default

## Browser Compatibility

This project should work in all modern browsers that support WebAssembly, including:
- Chrome/Chromium (version 57+)
- Firefox (version 53+)
- Safari (version 11+)
- Edge (version 16+)

## Next Steps and Extensions

Some ideas to extend this project:
- Add more mathematical operations (subtract, multiply, divide)
- Create more complex calculations in Go
- Improve error handling and input validation
- Add unit tests for the Go functions
- Add a loading indicator while the WebAssembly module loads

## Resources

- [Go WebAssembly Documentation](https://github.com/golang/go/wiki/WebAssembly)
- [MDN WebAssembly Documentation](https://developer.mozilla.org/en-US/docs/WebAssembly)