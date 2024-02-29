package main

import "fmt"

func main() {
	greeting("Test", "New York")
	bye()

	fmt.Println(myAge(2001))

	s := 3.0

	luasPersegi, kelilingPersegi := calculate(s)
	fmt.Printf("Luas Persegi: %v\n", luasPersegi)
	fmt.Printf("Keliling Persegi: %v\n", kelilingPersegi)

	// studentsList := []string{"john", "jane"}

	// students := myStudent(studentsList...)
	students := myStudent("john", "jane")

	fmt.Printf("students: %v",students)

	myKelas("golang-005", "john", "jane")


}

//Function represents
func greeting(name string, address string) {
	fmt.Printf("Hi! my name: %s, I live in %s\n", name, address)

}

//Function (Return)
func myAge(birthday int) string {
	age := 2024 - birthday
	thisYear := fmt.Sprintf("My age: %d", age)
	return thisYear
}

//Function with multiple (Return)
// func calculate(s float64) (float64, float64) {
// 	luas := s * s
// 	keliling := 4 * s

// 	return luas, keliling

// }
//Function (Predefined return value)
func calculate(s float64) (luas float64, keliling float64) {
	luas = s * s
	keliling = 4 * s

	return luas, keliling

}

//Function (Variadic function #1)

func myStudent(names ...string) []string {
	var absent []string

	for _, n := range names {
		absent = append(absent, n)
	}

	return absent
}


//Function (Variadic function #2)
func myKelas(className string, names  ...string) {
	var absent []string

	for _, n := range names {
		absent = append(absent, n)
	}

	fmt.Printf("ClassName: %s, and the name %v", className, absent)
}
func bye() {
	fmt.Println("bye")
}
