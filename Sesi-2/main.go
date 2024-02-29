package main

import "fmt"

func main() {

	// if condition dalam golang
	// warnaLampu := "merah"
	// fmt.Println("hasil if-else:")

	// if warnaLampu == "merah" {
	// 	fmt.Println("berhenti")
	// } else if warnaLampu == "kuning" {
	// 	fmt.Println("hati-hati")
	// } else if warnaLampu == "hijau" {
	// 	fmt.Println("hati-hati")
	// } else if warnaLampu == "ungu" {
	// 	fmt.Println("salto")
	// } else {
	// 	fmt.Println("tidur")
	// }

	// perulangan dalam golang
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("angka-i", i)
	// }

	// fmt.Println("============")

	// j := 0

	// for j < 3 {
	// 	fmt.Println("angka-j", j)
	// 	j++
	// }
	// fmt.Println("============")

	// var z = 0

	// for {

	// 	z++

	// 	if z == 1 || z == 2 {
	// 		fmt.Println("tidak ada")
	// 		continue
	// 	}
	// 	fmt.Println("angka-z", z)
	// 	if z > 20 {
	// 		break
	// 	}
	// }

	// nested loops
	// outerLoop:
	// 	for i := 0; i < 4; i++ {
	// 		fmt.Println("parent =", i)
	// 		for j := 0; j < 3; j++ {
	// 			fmt.Println("child =", j)

	// 			if j == 0 {
	// 				break outerLoop
	// 			}
	// 		}
	// 	}

	// array dalam golang
	var nilai int = 4
	_ = nilai

	// var nilaiRaport [4]int = [4]int{4, 5, 8, 9}

	// fmt.Printf("nilai rapot: %v\n", nilaiRaport)

	// for _, val := range nilaiRaport {
	// 	fmt.Printf("value: %d\n", val)
	// }

	var nilaiKelas [3][4]int = [3][4]int{{5, 5, 7, 4}, {6, 6, 7, 2}}

	fmt.Printf("nilai kelas: %v\n", nilaiKelas)

	for _, siswa := range nilaiKelas {
		fmt.Printf("siwa: %v\n", siswa)
		for _, nilai := range siswa {
			fmt.Println("nilai = ", nilai)
		}
	}

}
