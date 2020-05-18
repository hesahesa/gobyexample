// Di Go, _variabel_ dinyatakan secara eksplisit dan digunakan
// oleh kompilator untuk mengecek kecocokan tipe dari
// pemanggilan fungsi.

package main

import "fmt"

func main() {

	// `var` menyatakan 1 atau lebih variabel.
	var a = "initial"
	fmt.Println(a)

	// Kamu bisa menyatakan beberapa variabel
	// secara bersamaan.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go akan menyimpulkan tipe dari variabel yang
	// di-initialisasi.
	var d = true
	fmt.Println(d)

	// Variabel-variabel yang dinyatakan tanpa nilai
	// awal akan diberi _nilai nol_ (zero-valued).
	// Contohnya, nilai nol dari sebuah `int` adalah 0.
	var e int
	fmt.Println(e)

	// Sintaks `:=` adalah singkatan untuk menyatakan
	// dan memberi nilai awal sebuah variabel.
	// Contohnya di sini `var f string = "apel"`
	f := "apel"
	fmt.Println(f)
}
