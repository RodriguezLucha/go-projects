package main

import "fmt"

type contractInfo struct {
	zipCode int
	city    string
}

type person struct {
	firstName string
	lastName  string
	contractInfo
}

//Have to do a pointer otherwise it won't really update the name
func (personPtr *person) updateName(newName string) {
	(*personPtr).firstName = newName
}

func main() {
	fmt.Println("Start")

	rudy := person{
		firstName: "Rudy",
		lastName:  "Rodriguez",
		contractInfo: contractInfo{
			zipCode: 123456,
			city:    "Atlantis",
		},
	}

	rudy.updateName("Cactus")

	fmt.Printf("The person object is: %+v\n", rudy)
}
