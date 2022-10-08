package loops

import (
	"fmt"
)

//220 ve 284 arkadaş sayılardır.
//10 ve 65 arkadaş sayı değildir.
func Demo5() {
	sayi1 := 220
	sayi2 := 280
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 = toplam1 + i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 = toplam2 + i
		}
	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir", sayi1, sayi2)
	}

}
