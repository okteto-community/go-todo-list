package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type Todo struct {
	Task string `json:"task"`
	ID   string `json:"id"`
}

var db *sql.DB

func initDB() {
	var err error
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "todoapp")

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Create table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(36) PRIMARY KEY,
		task TEXT NOT NULL
	);`

	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Info("Database connected and initialized")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	todo := Todo{Task: task, ID: uuid.New().String()}
	
	_, err := db.Exec("INSERT INTO todos (id, task) VALUES ($1, $2)", todo.ID, todo.Task)
	if err != nil {
		log.Error("Failed to insert todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
	log.Info("saved todo item")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		log.Error("Failed to delete todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("Failed to get rows affected:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Info("deleted todo item")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, task FROM todos ORDER BY task DESC")
	if err != nil {
		log.Error("Failed to query todos:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
			log.Error("Failed to scan todo:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		log.Error("Row iteration error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
	log.Infof("got %d items", len(todos))
}

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz).Methods("GET")
	router.HandleFunc("/todo", getItems).Methods("GET")
	router.HandleFunc("/todo", createItem).Methods("POST")
	router.HandleFunc("/todo/{id}", deleteItem).Methods("DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	log.Info("Starting API server...")
	http.ListenAndServe(":8080", handler)
}
