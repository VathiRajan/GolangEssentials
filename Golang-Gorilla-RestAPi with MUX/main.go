package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var book []Book

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lasttname"`
}

//get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//i := r.GetBody()
	//Loop through the Books
	for _, item := range book {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//param := mux.Vars(r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(body)
	//body["author"]
	fmt.Println(bodyString)

	books := Book{}
	_ = json.NewDecoder(r.Body).Decode(&books)
	fmt.Println("Book One", books)

	//fmt.Println(json.NewDecoder(r.Body).Decode(&bookOne))
	books.ID = strconv.Itoa(rand.Intn(100000)) //Not Safe

	book = append(book, books)
	json.NewEncoder(w).Encode(books)

}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for index, value := range book {
		if value.ID == param["id"] {
			book = append(book[:index], book[index+1:]...)
			books := Book{}
			_ = json.NewDecoder(r.Body).Decode(&books)
			fmt.Println("Book One", books)

			//fmt.Println(json.NewDecoder(r.Body).Decode(&bookOne))
			books.ID = strconv.Itoa(rand.Intn(100000)) //Not Safe

			book = append(book, books)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(book)
}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for index, value := range book {
		if value.ID == param["id"] {
			book = append(book[:index], book[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(book)
}
func main() {
	//Init Router
	r := mux.NewRouter()

	//MockData - implemet DB
	book = append(book, Book{ID: "1",
		Isbn:  "44555",
		Title: "Java Dev", Author: &Author{
			Firstname: "John",
			Lastname:  "Corners",
		},
	})
	book = append(book, Book{ID: "2",
		Isbn:  "44455",
		Title: "Python Dev", Author: &Author{
			Firstname: "Mithran",
			Lastname:  "Corners",
		},
	})
	book = append(book, Book{ID: "3",
		Isbn:  "44555",
		Title: "GoLang Dev", Author: &Author{
			Firstname: "Vathi",
			Lastname:  "Corners",
		},
	})

	//Routing Handlers/Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)

}
