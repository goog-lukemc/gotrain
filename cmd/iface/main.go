package main

import (
	"fmt"
	"math"
)

func main() {
	area, _ := calculate(rect{
		width:  2,
		height: 2,
	})

	fmt.Printf("Rect:%v\n", area)

	area, _ = calculate(circle{
		radius: 10,
	})

	fmt.Printf("Circle:%v\n", area)

}

func calculate(g geometry) (float64, float64) {
	return g.area(), 0
}

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

// Exercise 15 minutes:
// Modify the code above and implement a function that calcutes the
// paremeter of a circle and rectangle

// Bonus:
// https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266

// Review the interface design in
// https://github.com/GoogleCloudPlatform/DIY-Tools/tree/master/gcp-data-drive

// Bonus: Take the quick lab at
// https://www.qwiklabs.com/focuses/10531?locale=id&parent=catalog
