package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getFriendByID(id int) (Friend, error) {
	friends := map[int]Friend{
		1: {
			ID:        1,
			Nama:      "Roihan Sori",
			Alamat:    "Kabupaten Bogor",
			Pekerjaan: "Fresh Graduate",
			Alasan:    " Bahasanya simpel, Mudah Dipahami, dan sedang banyak di cari perusahaan startup",
		},
		2: {
			ID:        2,
			Nama:      "Hafiz",
			Alamat:    "Surabaya",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin mengembangkan aplikasi web dengan cepat",
		},
		3: {
			ID:        3,
			Nama:      "Budi",
			Alamat:    "Bandung",
			Pekerjaan: "Developer",
			Alasan:    "Untuk memperluas kemampuan dalam membangun infrastruktur backend",
		},
		4: {
			ID:        4,
			Nama:      "Dinda",
			Alamat:    "Semarang",
			Pekerjaan: "Frontend Developer",
			Alasan:    "Ingin memperluas pengetahuan tentang pengembangan perangkat lunak terutama backend",
		},
		5: {
			ID:        5,
			Nama:      "Siti",
			Alamat:    "Yogyakarta",
			Pekerjaan: "UI/UX Designer",
			Alasan:    "Ingin memahami backend untuk meningkatkan kolaborasi dengan tim pengembang",
		},
		6: {
			ID:        6,
			Nama:      "Amelia",
			Alamat:    "Kabupaten Bogor",
			Pekerjaan: "Fresh Graduate",
			Alasan:    "Stabil dan Aman: Golang punya cara tangani kesalahan yang bagus, jadi aplikasi yang kamu buat cenderung lebih aman dan stabil.",
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
		fmt.Println(" ID belum dimasukkan! Gunakan: go run biodata.go 1")
		return
	}

	friendID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid ID ex: 1")
		return
	}

	friend, err := getFriendByID(friendID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan ikut kelas Golang: %s\n",
		friend.Nama,
		friend.Alamat,
		friend.Pekerjaan,
		friend.Alasan)
}
