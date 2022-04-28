package main
import "fmt"

type mayTinhLuong interface{
	tinhLuong() int
}
type daihan struct{
	luong int
	thuong int
}
type hopdong struct{
	luong int 
}
func (d daihan) tinhLuong() int{
	return d.luong + d.thuong
}
func (h hopdong) tinhLuong() int{
	return h.luong
}
func tongLuongNhanVien(m []mayTinhLuong) int{
	tong := 0 
	for _,v := range m{
		tong+= v.tinhLuong()
	}
	return tong
}
func main() {
	hd1 :=  hopdong{
		luong: 1000,
	}
	hd2 := hopdong{
		luong: 1500,
	}
	dh1 := daihan{
		luong: 2500,
		thuong: 500,
	}
	dh2 := daihan{
		luong: 3000,
		thuong: 1000,
	}
	nhanvien := []mayTinhLuong{hd1,hd2,dh1,dh2}
	fmt.Printf("tong luong cty la: %d $", tongLuongNhanVien(nhanvien))

}
