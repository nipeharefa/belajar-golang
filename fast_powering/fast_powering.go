// Link dari :
// https://github.com/trekhleb/javascript-algorithms/tree/master/src/algorithms/math/fast-powering
package main

import "fmt"

// Powering untuk mendapatkan hasil dari pangkat suatu bilangan
func Powering(base int, power int) int {
	// Jika power 0 maka return 1
	// Karna n^0 = 1
	if power == 0 {
		return 1
	}

	// Jika pangkat genap, maka power di bagi 2
	// Lalu dikalikan dengan hasilnya
	// Contoh power 2 di bagi 2, hasilnya jadi
	// (n*1) * (n*1)
	if power%2 == 0 {
		multipler := Powering(base, power/2)

		return multipler * multipler
	}

	multipler := Powering(base, power/2)

	return multipler * multipler * base
}

func main() {
	var base, power int

	fmt.Println("Input Base and Power : ")
	fmt.Scanf("%d %d", &base, &power)

	result := Powering(base, power)
	fmt.Println(result)
}
