package main 
import "fmt"
type class struct{
	place string
	size int
}
func main() {
	cls1 := class{"ban1", 12}
	cls2 := class{"ban2", 15}
	cls := class{
		place: "ban100",
		size: 12,
	}
	fmt.Println(cls)
	fmt.Println(cls1)
	fmt.Println(cls2)

}