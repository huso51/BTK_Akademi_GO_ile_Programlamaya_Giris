package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo3()
	//loops.Demo6()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.SelamVer("Engin")
	//functions.SelamVer("Derin")
	//var sonuc = functions.Topla(2, 6)
	//fmt.Println(sonuc * 10)

	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkarım : ", sonuc2)
	// fmt.Println("Çarpım : ", sonuc3)
	//fmt.Println("Bölüm : ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki sayı : ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayı : ", sayilar[0])

	//structs.Demo2()
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	//fmt.Println(error_handling.TahminEt2(99))

	//restful.Demo2()

	product, _ := project.AddProduct()
	fmt.Println(product)

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}
