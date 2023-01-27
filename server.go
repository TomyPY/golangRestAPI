package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type updateTask struct {
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	fmt.Fprintf(w, "Welcome to my API")
}

func getTasksRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByIdRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("ID read error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	log.Printf("ID doesn't exist")
	w.WriteHeader(400) // Return 400 Bad Request.
}

func createTaskRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	var newTask task
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal server error.
		fmt.Fprintf(w, "Insert a Valid Task")
		return
	}

	if err = json.Unmarshal(reqBody, &newTask); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func updateTaskRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(405)
		return
	}

	var newTask updateTask
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500)
		return
	}

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("ID read error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	if err = json.Unmarshal(reqBody, &newTask); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Name = newTask.Name
			tasks[i].Content = newTask.Content
			w.WriteHeader(200)
			return
		}
	}

	log.Printf("ID doesn't exist")
	w.WriteHeader(400)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		return
	}

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("ID read error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}

	log.Printf("ID doesn't exist")
	w.WriteHeader(400) // Return 400 Bad Request.
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasksRoute).Methods("GET")
	router.HandleFunc("/tasks", createTaskRoute).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTaskByIdRoute).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTaskRoute).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
