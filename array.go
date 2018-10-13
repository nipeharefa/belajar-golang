package main

import (
	"fmt"
)

func main() {
	var belumdiisi [2]string
	var nama = [2]string{"Nipe", "Setiawan"}

	belumdiisi[0] = "Indonesia"
	fmt.Printf("Panjang Array `nama` %d \n", len(nama))
	fmt.Printf("nama index 0 %s \n", nama[0])
	fmt.Printf("nama index 1 %s \n", nama[1])
	fmt.Printf("belumdiisi index ke 0 %s \n", belumdiisi[0])
}
