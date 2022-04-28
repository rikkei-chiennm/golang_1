package main
import "fmt"
func main() {
	emsalary := map[string]string{
		"chien":"hai",
		"long":"chin",
		"hoang":"nam",
	}
	check := "chie"
	value, ok := emsalary[check]
	if ok == true{
		fmt.Printf("key '%s' co trong map voi gia tri la '%s'", check, value)
	}else{
		fmt.Println("Nope")

	}
	for key, val := range emsalary{
		fmt.Printf("employee Salary[%s] = '%s' \n", key, val)
	}
	
}