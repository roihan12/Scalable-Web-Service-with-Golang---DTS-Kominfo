package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Roihan", Age: 22, Division: "IT"},
	{ID: 1, Name: "Roihan", Age: 22, Division: "IT"},
	{ID: 1, Name: "Roihan", Age: 22, Division: "IT"},
	{ID: 1, Name: "Roihan", Age: 22, Division: "IT"},
}

var PORT = ":3000"

func main() {
	http.HandleFunc("/", greeting)
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Aplication is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, world"

	fmt.Fprintln(w, msg)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
