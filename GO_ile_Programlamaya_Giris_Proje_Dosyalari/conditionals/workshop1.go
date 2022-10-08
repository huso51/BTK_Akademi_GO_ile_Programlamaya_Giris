package conditionals

import (
	"fmt"
)

func Demo3() {
	//üç adet int değişken tanımlayınız.
	//Ekrana en büyük olanı yazdırınız.

	var sayi1, sayi2, sayi3 int = 100, 1005, 18

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}

	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı : %v", enBuyuk)

}
