package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	Id      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(books)

	// database connection
	db := gormConnect()
	defer db.Close()

	posts := []Post{}
	db.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

//Get single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// database connection
	db := gormConnect()
	defer db.Close()

	post := Post{} // simple model
	db.Where("id = ?", params["id"]).Find(&post)
	// db.First(&post, "id= ?", params["id"])
	json.NewEncoder(w).Encode(post)

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// database connection
	db := gormConnect()
	defer db.Close()
	post := Post{}
	// var book Book
	_ = json.NewDecoder(r.Body).Decode(&post)
	db.Save(&post)
	json.NewEncoder(w).Encode(post)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// database connection
	db := gormConnect()
	defer db.Close()
	post := Post{}
	db.First(&post, params["id"]) // get value from database
	_ = json.NewDecoder(r.Body).Decode(&post)
	db.Save(&post)
	json.NewEncoder(w).Encode(post)

}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	// database connection
	db := gormConnect()
	defer db.Close()
	post := Post{}
	db.First(&post, params["id"]) // get value from database
	db.Delete(&post)
	json.NewEncoder(w).Encode(post)
}

//db connection
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "123456789"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "go-mysql-crud"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func routing() {
	//int Router
	r := mux.NewRouter()
	//Router Handler / Endpoints chill
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//listner
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	// database connection
	db := gormConnect()
	defer db.Close()
	//declare all routing
	routing()
}
