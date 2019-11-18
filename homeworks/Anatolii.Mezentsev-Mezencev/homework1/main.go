package main

import (
	"fmt"
	"math"
	"sort"
)

// Shape interface (model)
type Shape interface {
	Volume() float64
}

// Sphere struct (model)
type Sphere struct {
	radius float64
}

// Cone struct
type Cone struct {
	radius, height float64
}

// Rectangular struct
type Rectangular struct {
	side1, side2, height float64
}

// Volume for Sphere
func (c *Sphere) Volume() float64 {
	return 4.0 / 3.0 * math.Pi * math.Pow(c.radius, 3)
}

// Volume for Rectangular
func (c *Rectangular) Volume() float64 {
	return c.height * c.side1 * c.side2
}

// Volume for Cone
func (c *Cone) Volume() float64 {
	return c.height / 3 * math.Pi * math.Pow(c.radius, 2)
}

// List type
type List []Shape

func (s List) Len() int {
	return len(s)
}
func (s List) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s List) Less(i, j int) bool {
	return s[i].Volume() < s[j].Volume()
}

func describe(shapes List) {
	for _, s := range shapes {
		fmt.Printf("Shape: %#v = %v\n", s, s.Volume())
	}
}

func main() {
	m := make(List, 0, 4)
	m = append(m, &Sphere{10}, &Cone{5, 8}, &Rectangular{2, 3, 3})
	sort.Sort(m)
	describe(m)
}
