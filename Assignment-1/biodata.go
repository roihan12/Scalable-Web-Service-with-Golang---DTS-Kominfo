package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func getFriendByID(id int) (Friend, error) {
	friends := map[int]Friend{
		1: {
			ID:      1,
			Name:    "John Doe",
			Address: "123 Main St, Springfield, IL",
			Job:     "Software Engineer",
			Reason:  "Interested in learning Go",
		},
		2: {
			ID:      2,
			Name:    "Jane Smith",
			Address: "456 Oak Ave, Anytown, CA",
			Job:     "Data Scientist",
			Reason:  "Wants to expand programming skills",
		},
		3: {
			ID:      3,
			Name:    "Alice Johnson",
			Address: "789 Pine Rd, Somewhere, TX",
			Job:     "Web Developer",
			Reason:  "Heard good things about Go language",
		},
	}

	friend, exist := friends[id]
	if !exist {
		return Friend{}, fmt.Errorf("Friend with ID %d not found", id)
	}
	return friend, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <friendID>")
		return
	}

	friendID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid friend ID")
		return
	}

	friend, err := getFriendByID(friendID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan ikut kelas Golang: %s\n",
		friend.Name,
		friend.Address,
		friend.Job,
		friend.Reason)
}
