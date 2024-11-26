package main

import "fmt"

func main() {
	firstName := "dummyName"
	familyName := "dummyFamilyName"
	age := 50
	allergy := false

	fmt.Println(fmt.Sprintf("First name: %s", firstName))
	fmt.Println(fmt.Sprintf("Family name: %s", familyName))
	fmt.Println(fmt.Sprintf("Age: %d", age))
	fmt.Println(fmt.Sprintf("Allergy: %t", allergy))

}
