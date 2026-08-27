package main

import (
	"encoding/json"
	"net/http"
	"todo-go/src/controller"
)

func main() {
	taskController := &controller.TaskController{}
	taskController.CreateNewTask("Task1", "This is the first task")
	taskController.CreateNewTask("Task2", "This is the second task")
	taskController.CreateNewTask("Another Task", "One more task")
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

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method :("))
			return
		}

		tasks := taskController.Tasks()

		data, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "error al serializar las tareas", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	http.ListenAndServe(":3000", nil)
}
