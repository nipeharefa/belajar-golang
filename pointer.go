package main

import (
	"fmt"
)

func main() {
	// (*) atau disebut juga dengan asterisk
	// menandakan variabel bertipe pointer
	// ampersand(&) digunakan untuk mengambil nilai pointernya
	var angkaA int = 4
	var angkaB *int = &angkaA
	var angkaC int = angkaA
	// Pointer null
	var angkaD *int = nil

	fmt.Printf("Variabel angkaA disimpan di memory dengan alamat di %s \n", &angkaA)
	fmt.Printf("Angka => A : %d, B : %d, C : %d \n", angkaA, *angkaB, angkaC)

	angkaA = 1
	fmt.Printf("Variabel angkaA disimpan di memory dengan alamat di %s \n", &angkaA)
	fmt.Printf("Angka => A : %d, B : %d, C : %d \n", angkaA, *angkaB, angkaC)

	fmt.Printf("AngkaD disimpan di %s \n", angkaD)
	// Assign pointer angkaC ke D
	angkaD = &angkaC
	fmt.Printf("AngkaD disimpan di %s \n", *angkaD)

	// Lainnya
	// fmt.Printf("p \n", &1)

}
