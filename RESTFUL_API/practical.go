package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Marks []int  `json:"marks"`
}

var student_data []Student

func main() {
	fmt.Println("RESTAPI DUMMY")
	handlerMethods()
}

func handlerMethods() {
	Router := mux.NewRouter()
	Router.HandleFunc("/post", getAllStudents).Methods("GET")
	Router.HandleFunc("/post", addStudents).Methods("POST")
	Router.HandleFunc("/post/{id}", getSingleStudent).Methods("GET")
	Router.HandleFunc("/post/{id}", updateStudentData).Methods("PUT")
	Router.HandleFunc("/post/{id}", updateIndividualDataField).Methods("PATCH")
	Router.HandleFunc("/post/{id}", deleteStudentItem).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":5000", Router))
}
func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Student-Details", "Value")
	json.NewEncoder(w).Encode(student_data)
}
func addStudents(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	json.NewDecoder(r.Body).Decode(&newStudent)
	student_data = append(student_data, newStudent)
	w.Header().Set("Student-Details", "Value")
	json.NewEncoder(w).Encode(student_data)

}
func getSingleStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	parameter := id["id"]
	Id, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if Id >= len(student_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	students := student_data[Id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)

}

func updateStudentData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	parameter := id["id"]
	Id, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if Id >= len(student_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}

	var updatedData Student
	json.NewDecoder(r.Body).Decode(&updatedData)
	student_data[Id] = updatedData
	w.Header().Set("Student", "application/json")
	json.NewEncoder(w).Encode(updatedData)

}

func updateIndividualDataField(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	parameter := id["id"]
	Id, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if Id >= len(student_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}

	data := &student_data[Id]
	json.NewDecoder(r.Body).Decode(data)
	w.Header().Set("Student", "application/json")
	json.NewEncoder(w).Encode(data)

}

func deleteStudentItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	parameter := id["id"]
	Id, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if Id >= len(student_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	student_data = append(student_data[:Id], student_data[Id+1:]...)
	w.WriteHeader(200)
}
