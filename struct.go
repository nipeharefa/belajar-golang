package main

import (
	"fmt"
)

type User struct {
	id   int
	nama string
}

func main() {

	// Contoh Inisialisasi 1
	var a User
	a.id = 1
	a.nama = "Nipe Setiawan"

	fmt.Println("Struct 1")
	fmt.Println("========")
	fmt.Printf("Hai %s id anda adalah %d \n", a.nama, a.id)

	// Contoh Inisialisasi 2
	b := User{id: 1, nama: "Nipe Setiawan"}
	fmt.Println("Struct 2")
	fmt.Println("========")
	fmt.Printf("Hai %s id anda adalah %d \n", b.nama, b.id)

	// Contoh penggunaan pointer di struct
	var d *User
	d = &a
	d.nama = "Nipe"
	fmt.Println("Struct 2 Pointer")
	fmt.Println("========")
	fmt.Printf("Hai %s id anda adalah %d \n", a.nama, a.id)
}
