package main
import "fmt"
type student interface{
	read()
}
type teacher interface{
	write()
}
type People struct{}

func (p People) read() {
	fmt.Println("toi la hoc sinh")
}
func (p People) write(){
	fmt.Println("toi la thay co giao")
}
func main() {
	peo := People{}
	var stu student = peo
	stu.read()
	var tea teacher = peo
	tea.write()
}