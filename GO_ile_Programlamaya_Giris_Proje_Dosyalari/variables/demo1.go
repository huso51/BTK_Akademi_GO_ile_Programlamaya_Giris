package variables

import "fmt"

//camelCase
//PascalCase
func Demo1() {
	var metin string = "Merhaba Dünya!"
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	//integer
	var kdv int = -10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25.2
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	//true false
	durum = metin1 != metin2

	fmt.Println(durum)
	fmt.Println(!durum)
}
