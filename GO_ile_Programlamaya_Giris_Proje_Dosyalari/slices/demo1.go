package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "Engin"
	isimler[1] = "Derin"
	isimler[2] = "Salih"
	isimler = append(isimler, "Büşra")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
