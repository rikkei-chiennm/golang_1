// package main

// import (  
//     "fmt"
// )

// func number() int {  
//         num := 15 * 5 
//         return num
// }

// func main() {

//     switch num := number(); { //num is not a constant
//     case num < 50:
//         fmt.Printf("%d is lesser than 50\n", num)
//     case num < 100:
//         fmt.Printf("%d is lesser than 100\n", num)
//     case num < 200:
//         fmt.Printf("%d is lesser than 200", num)
//     }

// }



// package main

// import (  
//     "fmt"
// )

// func main() {  
//     fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
//     fruitslice := fruitarray[1:3]
//     fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
//     fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
//     fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
//     fmt.Print(fruitslice)
// }

package main

import (
	"fmt"
)

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    countriesCpy[1]="vietnam"
	return countries

}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
// hàm copy giúp cho khi chỉnh sửa slice không làm ảnh hưởng đến mảng chính 
