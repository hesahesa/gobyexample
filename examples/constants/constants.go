// Go mendukung _konstanta_ untuk karakter (character),
// deret karakter (string), boolean, dan nilai numerik.

package main

import (
	"fmt"
	"math"
)

// `const` mendeklarasikan sebuah nilai konstanta.
const s string = "konstanta"

func main() {
	fmt.Println(s)

	// Statemen `const` bisa muncul dimanapun statemen
	// `var` bisa.
	const n = 500000000

	// Ekspresi konstan melakukan aritmetika dengan
	// presisi arbitrari.
	const d = 3e20 / n
	fmt.Println(d)

	// Sebuah konstan numerik tidak punya tipe hingga
	// dia diberi sebuah tipe, misalnya dengan konversi
	// eksplisit.
	fmt.Println(int64(d))

	// Sebuah angka dapat diberikan sebuah tipe dengan
	// menggunakannya dalam konteks yang membutuhkan
	// tipe, misalnya pemberian nilai variabel (variable
	// assignment) atau pemanggilan fungsi.
	// Contohnya, disini `math.Sin` mengharapkan suatu `float64`.
	fmt.Println(math.Sin(n))
}
