package main

import (
	"fmt"

	"github.com/NicoBarragan/go-intermediate-course/pkg/composition"
	"github.com/NicoBarragan/go-intermediate-course/pkg/factory"
)

func main() {

	nicoEmployee := composition.AddBussinessEmployee(20, "Nico", "Dev", 1, 2)
	fmt.Println(nicoEmployee)

	// bad
	// nicoEmployee := BussinessEmployee{Person.Name: "Nico", Age: 20, Id: 0, YearsWorked: 2, Category: "Software Developer"}
	// fmt.Println(nicoEmployee)

	startWorking := nicoEmployee.StartWorking(true)
	endWorking := nicoEmployee.EndsWorking(true)

	fmt.Println("startWorking: ", startWorking)
	fmt.Println("endWorking: ", endWorking)

	fmt.Println("-------------")

	factory.Factory()
}
