package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"

	"github.com/rs/cors"
)

type Todo struct {
	Task string `json:"task"`
	ID   string `json:"id"`
}

var db *sql.DB

func initDB() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "todoapp"
	}

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	
	// Retry connection for up to 30 seconds
	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Errorf("Failed to connect to database: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		
		err = db.Ping()
		if err != nil {
			log.Errorf("Failed to ping database: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	
	if err != nil {
		log.Fatalf("Could not connect to database after 30 seconds: %v", err)
	}

	// Create table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(36) PRIMARY KEY,
		task TEXT NOT NULL
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	
	log.Info("Database connection established and table created")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	todo := Todo{Task: task, ID: uuid.New().String()}
	
	_, err := db.Exec("INSERT INTO todos (id, task) VALUES ($1, $2)", todo.ID, todo.Task)
	if err != nil {
		log.Errorf("Failed to insert todo: %v", err)
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
		log.Errorf("Failed to delete todo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Failed to get rows affected: %v", err)
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
	rows, err := db.Query("SELECT id, task FROM todos")
	if err != nil {
		log.Errorf("Failed to query todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var all []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Task)
		if err != nil {
			log.Errorf("Failed to scan todo: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		all = append(all, todo)
	}
	
	if err = rows.Err(); err != nil {
		log.Errorf("Row iteration error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].Task > all[j].Task
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all)
	log.Infof("got %d items", len(all))
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
