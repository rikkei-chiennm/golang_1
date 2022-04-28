package main
import "fmt"
type student struct{
	id string
	class string
}
func main() {
	stu1 := student{
		id: "B19DCCN100",
		class: "CN04",
	}
	stu2 := student{
		id: "B19DCCN101",
		class: "CN05",
	}
	stu3 := student{
		id: "B19DCCN005",
		class: "CN01",
	}
	studentInfor := map[string]student{
		"chien" : stu1,
		"manh" : stu2,
		"Long" : stu3,
	}
	for key, val := range studentInfor{
		fmt.Printf("%s có mã sinh viên %s lop %s \n", key, val.id, val.class)
	}
}