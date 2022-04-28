package main
import "fmt"

type Employee struct{
	name string
	salary int
}
func (e Employee) disEmployee(){
	fmt.Printf("Salary of %s is %d", e.name, e.salary)
}
func main() {
emp1 := Employee{
	name: "chien",
	salary: 9,
}
emp1.disEmployee()
	
}