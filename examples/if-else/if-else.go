// Percabangan dengan `if` dan `else` di Go
// cukuplah jelas.

package main

import "fmt"

func main() {

	// Ini adalah sebuah contoh sederhana.
	if 7%2 == 0 {
		fmt.Println("7 genap")
	} else {
		fmt.Println("7 ganjil")
	}

	// Kamu bisa punya sebuah `if` tanpa sebuah `else`
	if 8%4 == 0 {
		fmt.Println("8 habis dibagi 4")
	}

	// Sebuah statemen dapat mendahului kondisional; variabel
	// apapun yang dinyatakan dalam statemen ini akan
	// tersedia di semua percabangan.
	if num := 9; num < 0 {
		fmt.Println(num, "negatif")
	} else if num < 10 {
		fmt.Println(num, "punya 1 digit")
	} else {
		fmt.Println(num, "punya banyak digit")
	}
}

// Catatan: kamu tidak perlu tanda kurung "()" disekitar kondisi
// di Go, tapi tanda kurung kurawal "{}" diperlukan.
