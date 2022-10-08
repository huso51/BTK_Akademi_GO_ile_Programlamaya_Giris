package structs

import (
	"fmt"
)

func Demo1() {
	x:= product{"Bilgisayar", 500, "XYZ"}
	fmt.Println(x)
	fmt.Println(product{"Bilgisayar", 500, "ABC"})
	fmt.Println(product{name: "Bilgisayar"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
