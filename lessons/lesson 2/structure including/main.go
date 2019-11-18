package main

import "fmt"

// Printable -
type Printable interface {
	Print() string
}

// Person -
type Person struct {
	Firstname string
	Lastname  string
}

func (p *Person) Print() string {
	return p.Firstname + " " + p.Lastname
}

// Employee -
type Employee struct {
	*Person
	Job string
}

func (e *Employee) Print() string {
	return e.Person.Print() + " as " + e.Job
}

func printSomthing(i Printable) {
	fmt.Printf("(%v, %T) = %s\n", i, i, i.Print())
}

func printWithTypeSelection(i Printable) {
	switch v := i.(type) {
	case *Person:
		fmt.Printf("Person: %s\n", v.Firstname)
	default:
		fmt.Printf("%s\n", i.Print())
	}
}

func printWithTypeSelectionAndAnnonimInterface(i interface{}) {
	switch v := i.(type) {
	case *Person:
		fmt.Printf("Person: %s\n", v.Firstname)
	case *Employee:
		fmt.Printf("%s\n", v.Print())
	case string:
		fmt.Printf("%s\n", v)
	default:
		fmt.Printf("We can't print this type (%T) %v\n", v, v)
	}
}

func main() {
	p := &Person{
		Firstname: "John",
		Lastname:  "Smith",
	}
	printSomthing(p)

	printWithTypeSelection(p)
	printWithTypeSelectionAndAnnonimInterface(p)

	e := &Employee{
		Job: "Developer",
	}
	e.Person = p

	printSomthing(e)

	var i interface{}

	i = e

	s, ok := i.(string)
	if ok {
		fmt.Printf("OK %s\n", s)
	} else {
		fmt.Printf("Sorry\n")
	}

	fmt.Printf("%v, %T\n", s, s)

	printWithTypeSelection(e)
	printWithTypeSelectionAndAnnonimInterface(e)
	printWithTypeSelectionAndAnnonimInterface("Some string value")
	printWithTypeSelectionAndAnnonimInterface([]int{0, 2, 3})
	printWithTypeSelectionAndAnnonimInterface(34.5)

}
