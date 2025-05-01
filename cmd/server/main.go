package main

import (
	"encoding/json"
	"fmt"
	Gmux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func MakeWebHandler() http.Handler {
	mux := Gmux.NewRouter()
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentListHandler).Methods("GET")
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 18, 87}
	students[2] = Student{2, "bbb", 19, 98}
	lastId = 2

	return mux
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	vars := Gmux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	port := ":3030"
	fmt.Printf("WebServer Started %s\n", port)
	http.ListenAndServe(port, MakeWebHandler())
}
