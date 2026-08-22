package main

import "net/http"

func main() {
	http.Handle("/style.css", http.FileServer(http.Dir("./public")))
	http.Handle("/js/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./public/index.html")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method :( \nOnly GET method accepted"))
		}
	})

	http.HandleFunc("/addtask", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./public/addtask.html")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method :( \nOnly GET method accepted"))
		}
	})

	http.ListenAndServe(":3000", nil)
}
