package main

import "syscall/js"

// Add is our function that adds two numbers
func Add(this js.Value, args []js.Value) interface{} {
	// Get the two numbers from the arguments
	a := args[0].Int()
	b := args[1].Int()

	// Return the sum
	return js.ValueOf(a + b)
}

func main() {
	// Create a channel to keep the program running
	c := make(chan struct{}, 0)

	// Register our Add function to be callable from JavaScript
	js.Global().Set("addNumbers", js.FuncOf(Add))

	// Log a message to confirm the WASM module has loaded
	println("WebAssembly module loaded")

	// Keep the program running
	<-c
}
