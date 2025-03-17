package main

import "fmt"

type Dog struct {
	name string
}

func (d Dog) Name() string {
	fmt.Println("dog name")
	return ""
}
func (d *Dog) SetName(name string) {
	fmt.Println("SetName", name)
}

type Pet interface {
	Name() string
	SetName(name string)
}

func main() {
	var dog1 *Dog
	fmt.Println("The first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("The second dog is nil. [wrap1]")
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
		fmt.Printf("The pet type is %T\n", pet)
	}
	var dog3 *Dog = nil
	pet = dog3
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
		fmt.Printf("%T\n", pet)
	}
	//fmt.Println(pet.Name())
	pet.SetName("233")

}
