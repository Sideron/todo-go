package main

import "net/http"

func main() {
	// 1. Serve static files (like style.css) from the "public" folder
	// This maps requests like "/style.css" to "./public/style.css"
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	/*http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./public/main.html")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method :( \nOnly GET method accepted"))
		}
	})*/

	http.ListenAndServe(":3000", nil)
}
