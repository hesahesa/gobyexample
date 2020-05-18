// Go punya beberapa tipe nilai (value): deret
// karakter (string), bilangan bulat (integer),
// bilangan titik ambang (float), boolean, dll.
// Ini adalah beberapa contoh-contoh dasar.

package main

import "fmt"

func main() {

	// String, bisa digabung dengan menggunakan `+`.
	fmt.Println("go" + "lang")

	// Integer dan float.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean, dengan operator boolean seperti
	// yang kamu harapkan.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
