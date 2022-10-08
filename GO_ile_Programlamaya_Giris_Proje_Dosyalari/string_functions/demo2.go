package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasSuffix(isim, "en"))
	fmt.Println(s.Index(isim, "gi"))
	harfler := []string{"e", "n", "g", "i", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 5))
}
