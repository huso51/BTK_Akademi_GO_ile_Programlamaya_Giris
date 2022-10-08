package defer_statement

import (
	"fmt"
)

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	p = product{productName: "Mouse", unitPrice: 45}

	fmt.Println("İşlem başarılı")
	fmt.Println(p.productName)
	defer p.save()
}
