// Go mendukung
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>fungsi rekursif</em></a>.
// Ini adalah contoh klasik faktorial.

package main

import "fmt"

// Fungsi `fact` ini memanggil dirinya sendiri hingga dia
// sampai ke kasus dasar (base case), yaitu `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
