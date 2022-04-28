package main
import (
	"fmt"
)
func main() {
	slice1 := [] string {"chien", "thang", "bat", "tan"}
	slice2 := slice1[1:]
	var slice3 [] string
	
	slice3 = append(slice3, slice1...)
	slice3[1]="hoang"

	fmt.Print(slice3,"\n", slice1,"\n", slice2)
}
//khi truy cập hàm vào thì slice bị thay đổi, còn arrays thì không ảnh hưởng
// sao chép slice này qua một slice khác bằng câu lệnh dòng 8, như vậy thì slice 3 thay đổi thì slice 2 và 1 không bị thay dổi 

// length 
// cap
// slic1 lenth 4 cap 4
// slice append 1 ptu lenth 5 cap 8 