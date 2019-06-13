package main

import (
	"encoding/json"
	"fmt"
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

//Get all Posts
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(Posts)

	// database connection
	db := gormConnect()
	defer db.Close()

	posts := []Post{}
	db.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

//Get single Post
func getPost(w http.ResponseWriter, r *http.Request) {
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

//Create a new Post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// database connection
	db := gormConnect()
	defer db.Close()
	post := Post{}
	// var Post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	db.Save(&post)
	json.NewEncoder(w).Encode(post)

}

func updatePost(w http.ResponseWriter, r *http.Request) {
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
func deletePost(w http.ResponseWriter, r *http.Request) {
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
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/api/posts", createPost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	//listner
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	// database connection
	db := gormConnect()
	defer db.Close()

	fmt.Println("app started ....")
	//declare all routing
	routing()
}
