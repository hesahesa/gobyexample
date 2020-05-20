// _Gorutin_ (Goroutine) adalah sebuah thread eksekusi
// yang ringan.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Misalkan kita punya sebuah pemanggilan fungsi `f(s)`.
	// Ini bagaimana kita memanggilnya dengan cara normal,
	// menjalankan fungsi tersebut secara sinkronis.
	f("langsung")

	// Untuk menjalankan fungsi ini di sebuah gorutin, gunakan
	// `go f(s)`. Gorutin ini akan dieksekusi secara bersamaan
	// dengan si pemanggil.
	go f("gorutin")

	// Kamu juga bisa memulai sebuah gorutin untuk pemanggilan
	// fungsi anonim.
	go func(msg string) {
		fmt.Println(msg)
	}("anonim")

	// Dua buah pemanggilan fungsi kita berjalan secara bersamaan
	// dalam gorutin yang terpisah. Tunggu sampai mereka selesai
	// (untuk pendekatan yang lebih kuat, gunakan [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("selesai")
}
