package main

//import "fmt"
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Todo represents a single todo item
type Todo struct {
	ID    string
	Title string
	Done  bool
}

// In-memory storage - slice of todos
//20260519 removal due to inclusion of PostgreSQL
// var todos []Todo
// var nextID int = 1

// Declaring global db variable
var db *sql.DB

// 20260514 Lesson: Adding GET HANDLER
// Get /todos -returns all todos from PostgreSQL
func getTodos(w http.ResponseWriter, r *http.Request) {
	//20260519 enhancement in adding PostgreSQL
	rows, err := db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// 20260514 Lesson: Adding POST handler
func createTodo(w http.ResponseWriter, r *http.Request) {
	// var todo Todo
	// json.NewDecoder(r.Body).Decode(&todo) //&todo = pointer to todo so Decode can modify the original //Converts JSON > Go Struct
	// todo.ID = fmt.Sprintf("%d", nextID)
	// nextID++
	// todos = append(todos, todo)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(todo)

	//20260519 enhancement to use PostgreSQL
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := db.QueryRow(
		"INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id",
		todo.Title, todo.Done,
	).Scan(&todo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

// 20260514 Lesson: Adding Delete Handler
// DELETE /todos/{id} - deletes a todo by ID
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Path[len("/todos/"):]
	// for i, todo := range todos {
	// 	if todo.ID == id {
	// 		todos = append(todos[:i], todos[i+1:]...)
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	// }
	// http.Error(w, "Todo not found!", http.StatusNotFound)

	//20260519 enhancement to use PostgreSQL for delete handler
	id := r.URL.Path[len("/todos/"):]
	result, err := db.Exec("DELETE FROM todos WHERE id =$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo not found!!!", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// 20260514 Lesson: Adding Put Handler
// PUT/todos/{id} - updates a todo by ID
func updateTodo(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Path[len("/todos/update/"):]
	// var updatedTodo Todo
	// json.NewDecoder(r.Body).Decode(&updatedTodo)
	// for i, todo := range todos {
	// 	if todo.ID == id { //if match found using ID property of request params {id} to []todos struct
	// 		updatedTodo.ID = id
	// 		todos[i] = updatedTodo
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(updatedTodo)
	// 		return
	// 	}
	// }
	// http.Error(w, "Todo not found to update!!!", http.StatusNotFound)

	//20260519 enhancement to include PostgreSQL for Update handler
	id := r.URL.Path[len("/todos/update/"):]
	var updatedTodo Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	result, err := db.Exec(
		"UPDATE todos SET title = $1, done = $2 WHERE id = $3",
		updatedTodo.Title, updatedTodo.Done, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo not found!!!", http.StatusNotFound)
		return
	}
	updatedTodo.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)

}

// 20260519 adding DB initialization method
func initDB() {
	connStr := "host=localhost port=5432 user=postgres password=<YOUR_PASSWORD> dbname=tododb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
}

func main() {
	//20260519 init db call
	initDB()
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
