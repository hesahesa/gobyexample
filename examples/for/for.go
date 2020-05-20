// `for` adalah satu-satunya fitur perulangan di Go.
// Ini adalah beberapa tipe dasar perulangan `for`.

package main

import "fmt"

func main() {

	// Tipa yang paling dasar, dengan sebuah kondisi.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Perulangan `for` dengan nilai awal/kondisi/setelah.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` tanpa sebuah kondisi akan berulang terus-menerus
	// sampai kamu `break` keluar dari perulangan atau
	// `return` dari blok fungsi.
	for {
		fmt.Println("loop")
		break
	}

	// Kamu juga bisa `continue` untuk melanjutkan ke
	// iterasi selanjutnya dari perulangan.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
