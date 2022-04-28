package main 
import "fmt"

type player struct{
	number int
	size string
}

func main() {
	play1  := player{
		number: 10,
		size: "L",
	}
	play2 := player{
		number: 19,
		size: "XL",
	}
	play3 := player{
		number: 21,
		size: "M",
	}
	playInfor := map[string]player{
		"Chien" : play1,
		"Hai" : play2,
		"Long" : play3,
	}
	for key, val := range playInfor{
		fmt.Printf("Cau thu '%s' mang ao so '%d' size '%s' \n", key, val.number, val.size)
	}
}