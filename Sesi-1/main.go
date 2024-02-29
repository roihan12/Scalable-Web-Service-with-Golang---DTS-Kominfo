package main

import "fmt"

func main() {
	//Komentar
	fmt.Println("Hello, world!")

	//Variabel
	//Variabel with type data
	// var name string = "Roihan Sori"
	// var age int = 22

	// Variable without data type(Short Declaration)
	contoh := "Contoh short declaration"
	_ = contoh

	//Multiple variable declarations

	var students1, students2, students3 string = "Student1", "Student2", "Student3"
	// var first, second, third int

	first, second, third := "1", 2, "4"
	fmt.Println(students1, students2, students3)

	fmt.Println(first, second, third)



	var nama = "Roihan"
	var age = 22
	var address = "Jl.cemara"

	fmt.Printf("Halo nama ku %s, umur ku aku adalah %d, dan aku tinggal di %s", nama, age, address)

}
