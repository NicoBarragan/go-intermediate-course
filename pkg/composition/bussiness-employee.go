package composition

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int8
}

type Bussiness struct {
	Category string
}

type Employee struct {
	Id          int64
	YearsWorked int8
}

type BussinessEmployee struct {
	// Anonymous struct types: without instanciating the Composition
	Person
	Employee
	Bussiness
	// For ex, this won't be anonymous:
	// p Person
}

func AddBussinessEmployee(age int8, name string, category string, id int64, yearsWorked int8) *BussinessEmployee {
	// It could also be done by this way
	// newEmployee := BussinessEmployee{
	// 	Person:    Person{Name: name},
	// 	Employee:  Employee{},
	// 	Bussiness: Bussiness{},
	// }

	// But this way is only possible when the struct types are anonymous:
	newEmployee := &BussinessEmployee{}
	newEmployee.Age = age
	newEmployee.Name = name
	newEmployee.Category = category
	newEmployee.Id = id
	newEmployee.YearsWorked = yearsWorked

	return newEmployee
}

func (b *BussinessEmployee) StartWorking(started bool) string {
	if started {
		timeStarted := time.Now().String()
		return fmt.Sprintf("Started working at: %s hs", timeStarted)
	}

	return "Is not working right now"
}

func (b *BussinessEmployee) EndsWorking(ended bool) string {
	if ended {
		timeStarted := time.Now().String()
		return fmt.Sprintf("Ended working at: %s hs", timeStarted)
	}

	return "It has not ended working yet"
}
