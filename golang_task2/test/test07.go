package test

import "fmt"

type Person struct {
	Name string
	Age  uint
}

type Employee struct {
	EmployeeID uint
	Person
}

func (emp Employee) PrintInfo() {
	fmt.Printf("员工工号：%v,员工姓名：%v,员工年龄：%v", emp.EmployeeID, emp.Name, emp.Age)
}

func Test07() {
	emp := Employee{
		EmployeeID: 20,
		Person: Person{
			Name: "张三",
			Age:  28,
		},
	}

	emp.PrintInfo()
}
