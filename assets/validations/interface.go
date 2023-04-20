package main

import (
	"fmt"
	"sort"
)

type Animal interface {
	MakeSound() string
}

type Dog struct {
	Name string
}

func (d *Dog) MakeSound() string {
	return "Jhow"
}

func AnimalDound(a Animal) {
	fmt.Println(a.MakeSound())
}

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	people := []Person{
		{"Alice", 25, "New York"},
		{"Bob", 30, "Los Angeles"},
		{"Charlie", 35, "Chicago"},
		{"David", 40, "New York"},
	}

	filtered := []Person{}
	for _, p := range people {
		if p.City == "New York" {
			filtered = append(filtered, p)
		}
	}

	fmt.Println(filtered)

	nums := []int{5, 2, 8, 1, 6, 9, 3, 7, 4}

	sort.Ints(nums)

	fmt.Println(nums)

	names := []string{"Zelda", "Flavio", "Alice", "Charlie", "Bob", "David", "Ana"}

	sort.Strings(names)

	fmt.Println(names)
}
