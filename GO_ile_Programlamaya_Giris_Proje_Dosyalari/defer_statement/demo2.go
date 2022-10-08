package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("çift sayı çalıştı")
		return "Çift sayıdır"
	}

	if sayi%2 != 0 {
		return "Tek sayıdır"
	}

	return "Belli değil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
