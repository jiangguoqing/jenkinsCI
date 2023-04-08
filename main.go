package main

import (
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from littlejgq,v2"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet...v3"))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet...v3"))
}


func main() {
	// Register the two new handler functions and corresponding URL patterns wi
	// the servemux, in exactly the same way that we did before.

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	//apollo动态获取配置。
	//日志处理。
	log.Fatal(err)
}
