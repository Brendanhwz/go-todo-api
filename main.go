package main

//import "fmt"
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Todo represents a single todo item
type Todo struct {
	ID    string
	Title string
	Done  bool
}

// In-memory storage - slice of todos
var todos []Todo
var nextID int = 1

// 20260514 Lesson: Adding GET HANDLER
// Get /todos -returns all todos as JSON
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// 20260514 Lesson: Adding POST handler
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo) //&todo = pointer to todo so Decode can modify the original //Converts JSON > Go Struct
	todo.ID = fmt.Sprintf("%d", nextID)
	nextID++
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

// 20260514 Lesson: Adding Delete Handler
// DELETE /todos/{id} - deletes a todo by ID
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/"):]
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found!", http.StatusNotFound)
}

// 20260514 Lesson: Adding Put Handler
// PUT/todos/{id} - updates a todo by ID
func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/update/"):]
	var updatedTodo Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	for i, todo := range todos {
		if todo.ID == id { //if match found using ID property of request params {id} to []todos struct
			updatedTodo.ID = id
			todos[i] = updatedTodo
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	http.Error(w, "Todo not found to update!!!", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/todos", getTodos)
	http.HandleFunc("/todos/create", createTodo)
	//Adding Delete controller to main
	http.HandleFunc("/todos/", deleteTodo)
	//Adding Put/Update controller to main
	http.HandleFunc("/todos/update/", updateTodo)
	fmt.Println("TODO API starting on port 8080...")
	http.ListenAndServe(":8080", nil)
	// fmt.Println("TODO API starting...")
}
