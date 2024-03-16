package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	students = append(students, &Student{Id: "s001", Name: "Bourne", Grade: 4})
	students = append(students, &Student{Id: "s002", Name: "John", Grade: 5})
	students = append(students, &Student{Id: "s003", Name: "Jane", Grade: 6})
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)

	server.Addr = ":9000"

	fmt.Println("Server listening on localhost", server.Addr)
	server.ListenAndServe()

}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}

	if !AllowOnlyGet(w, r) {
		return
	}

	id := r.URL.Query().Get("id")
	fmt.Println(id)
	if id := r.URL.Query().Get("id"); id != "" {

		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
