package main
import "fmt"
func main() {
	a:=[3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%d ", v2)
        }
        fmt.Printf("\n")
    }
}