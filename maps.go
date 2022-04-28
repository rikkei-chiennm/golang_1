// package main
// import "fmt"
// func main() {
// 	bando := make(map[string]int)
// 	bando["vietnam"]= 90
// 	bando["USA"]=900
// 	bando["thailand"]=65
// 	fmt.Println(bando)
// 	for key, value :=range bando{
// 		fmt.Printf("%s has area of %d \n", key, value)
// 	}
// }
 package main
 import "fmt"
 func main() {
	 bando := map[string]int{
		 "vietnam" : 90,
		 "USA" : 900,
		 "thailand": 65,
	 }
	 bando["korea"]=48
	for key, value :=range bando{
		fmt.Printf("%s has area of %d \n", key, value)
	}
 }