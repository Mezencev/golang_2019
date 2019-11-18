package main

import (
	"fmt"
	"math"
	"sort"
)

// Shape3D - a three-dimensional geometric shape
type Shape3D interface {
	Volume() float64
}

// Sphere - a perfectly round geometrical object in three-dimensional space
type Sphere struct {
	Radius float64
}

// Volume of sphere
func (s *Sphere) Volume() float64 {
	return 4. / 3 * math.Pi * math.Pow(s.Radius, 3)
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere: { radius: %v }", s.Radius)
}

// Cone - a three-dimensional geometric shape that tapers smoothly from a flat base to a point called the apex or vertex.
type Cone struct {
	Radius float64
	Height float64
}

// Volume of cone
func (c *Cone) Volume() float64 {
	return c.Height / 3 * math.Pi * c.Radius * c.Radius
}

func (c *Cone) String() string {
	return fmt.Sprintf("Cone: { radius: %v, height: %v }", c.Radius, c.Height)
}

// Cube - the parallelepiped with Oh symmetry, which has six congruent square faces.
type Cube struct {
	Side float64
}

// Volume of cube
func (c *Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (c *Cube) String() string {
	return fmt.Sprintf("Cube: { side: %v }", c.Side)
}

// Shapes implements sort.Interface for []Shape3D based on shape volume
type Shapes []Shape3D

func (s Shapes) Len() int {
	return len(s)
}
func (s Shapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Shapes) Less(i, j int) bool {
	return s[i].Volume() < s[j].Volume()
}

func (s Shapes) appendShape(shape Shape3D) {
	fmt.Println("Appending 3D shape:", shape)
	s = append(s, shape)
}

func printShapeVolumes(shapes Shapes) {
	fmt.Println("List of shapes:")
	for i, s := range shapes {
		fmt.Printf(" Shape index: %v; Volume: %v - %v\n", i, s.Volume(), s)
	}
}

func sortShapes(shapes Shapes) Shapes {
	fmt.Printf("\nUnsorted ")
	printShapeVolumes(shapes)

	sort.Sort(shapes)

	fmt.Printf("\nSorted ")
	printShapeVolumes(shapes)
	return shapes
}

func main() {
	shapes := make(Shapes, 0, 4)

	shapes.appendShape(&Sphere{5})
	shapes = appendShape(shapes, &Sphere{6.2})
	shapes = appendShape(shapes, &Cube{10})
	shapes = appendShape(shapes, &Cone{10, 5})

	shapes = sortShapes(shapes)
}
