package string_functions

//alias
import (
	"fmt"
	s "strings"
)

//case sensitive
//ascii
func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "a"))
	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
