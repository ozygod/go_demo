package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name     string
	category string
}

func (d *Dog) SetName(name string) {
	d.name = name
}
func (d Dog) Name() string {
	return d.name
}
func (d Dog) Category() string {
	return d.category
}

func main() {
	d := Dog{name: "dog"}
	var pet Pet = &d
	d.SetName("cat")
	fmt.Println(pet.Name())
	fmt.Println(d.Name())
}
