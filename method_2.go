package main

import (  
    "fmt"
)

// type Rectangle struct {  
//     length int
//     width  int
// }

// type Circle struct {  
//     radius float64
// }

// func (r Rectangle) Area() int {  
//     return r.length * r.width
// }

// func (c Circle) Area() float64 {  
//     return math.Pi * c.radius * c.radius
// }

// func main() {  
//     r := Rectangle{
//         length: 10,
//         width:  5,
//     }
//     fmt.Printf("Area of rectangle %d\n", r.Area())
//     c := Circle{
//         radius: 12,
//     }
//     fmt.Printf("Area of circle %f", c.Area())
// }
type hinhvuong struct{
	canh int
}
type hinhbh struct{
	duongcao int
	canhday int
}
func (h hinhvuong) Areah() int  {
	return h.canh*h.canh
}
func (b hinhbh) Areab() int {
	return b.canhday*b.duongcao
}
func main() {
	hv := hinhvuong{
		canh: 5,
	}
	hbh := hinhbh{
		canhday: 15,
		duongcao: 10,
	}
	fmt.Print(hbh.Areab())
	fmt.Println()
	fmt.Print(hv.Areah())
}
