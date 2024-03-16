package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Response struct {
	Status Status `json:"status"`
}

func generateRandomData() Status {
	return Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}
}

func updateJSONFile(filename string) {
	for {
		data := generateRandomData()
		file, err := json.MarshalIndent(Response{Status: data}, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		err = WriteToFile(filename, file)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Printf("Updated JSON file: %s\n", filename)
		time.Sleep(15 * time.Second)
	}
}

func WriteToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}



func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET supported.", http.StatusNotFound)
		return
	}
	data := generateRandomData()
	response := Response{
		Status: data,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	filename := "data.json"
	go updateJSONFile(filename)

	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}
