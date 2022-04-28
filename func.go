package main

import (
	"fmt"
	"math"
)

func tamgiac(c1, c2, c3 int) (int, float64) {
	chuvi := c1 + c2 + c3
	var diendich float64 = math.Sqrt(float64(chuvi*(chuvi-c1)*(chuvi-c2)*(chuvi-c3)))
	return chuvi, diendich
}
func main() {
	c1, c2, c3 := 2, 3, 4
	chuvi, dientich := tamgiac(c1, c2, c3)
	fmt.Printf("dien tich tam giac la %.2f va chu vi la %d", dientich, chuvi)
}
// muon chỉ trả về dientich thì dùng dientich, _ := tamgiac(c1, c2, c3)
// slice [] 
// array [...]