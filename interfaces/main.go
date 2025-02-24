package main

import "fmt"

type Person interface {
	getName() string
	getAge() int
	getSalary() float64
}

type Employee struct {
	name   string
	age    int
	salary float64
}

type Customer struct {
	name    string
	age     int
	balance float64
}

type Director struct {
	name   string
	age    int
	salary float64
	bonus  float64
}

func (e Employee) getName() string {
	return e.name
}

func (e Employee) getAge() int {
	return e.age
}

func (e Employee) getSalary() float64 {
	return e.salary
}

func (e Director) getSalary() float64 {
	return e.salary + e.bonus
}

func (d Director) getAge() int {
	return d.age
}

func (d Director) getName() string {
	return d.name
}

func (c Customer) getName() string {
	return c.name
}

func (c Customer) getAge() int {
	return c.age
}

func (c Customer) getSalary() float64 {
	return c.balance
}

func printDetails(p Person) {
	switch p.(type) {
	case Employee:
		fmt.Println("Employee")
	case Customer:
		fmt.Println("Customer")
	case Director:
		fmt.Println("Director")
	}

	fmt.Println("Name:", p.getName())
	fmt.Println("Age:", p.getAge())
	fmt.Println("Salary:", p.getSalary())
	fmt.Println()
}

func main() {
	emp := Employee{"John", 30, 1000.0}
	customer := Customer{"Mary", 25, 500.0}
	director := Director{"Jane", 40, 2000.0, 500.0}

	printDetails(emp)
	printDetails(customer)
	printDetails(director)
}
