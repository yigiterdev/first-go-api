package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type Person struct {
	Name	string	`json:"name"`
	Surname	string	`json:"surname"`
	Age		int		`json:"age"`
}
var people []Person
func getPeople(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(people)
}
func addPeople(w http.ResponseWriter,router *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var newPerson Person
	_ = json.NewDecoder(router.Body).Decode(&newPerson)
	people = append(people,newPerson)
	json.NewEncoder(w).Encode(newPerson)
}
func main() {
	router := mux.NewRouter()
	//Mock Data
	people = append(people,Person{Name:"Ahmet Buğra",Surname:"Yiğiter",Age:21})
	people = append(people,Person{Name:"Buğra",Surname:"Şenocak",Age:22})

	router.HandleFunc("/api/people",getPeople).Methods("GET")
	router.HandleFunc("/api/people",addPeople).Methods("POST")
	http.ListenAndServe(":8000",router)
	log.Fatal(http.ListenAndServe(":8000",router))
}