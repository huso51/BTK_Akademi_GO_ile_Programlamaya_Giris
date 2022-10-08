package loops

import (
	"fmt"
)

//1. Kullanıcıdan bir sayı girmesini isteyiniz.
//23 : Asaldır
func Demo4() {
	sayi := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("Asaldır")
	} else {
		fmt.Println("Asal değildir")
	}

}
