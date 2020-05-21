// _Kanal_ (Channel) adalah sebuah pipa yang menghubunngkan
// beberapa gorutin. Kamu bisa mengirimkan nilai ke kanal
// dari satu gorutin dan menerima nilai tersebut di
// gorutin lainnya.

package main

import "fmt"

func main() {

	// Buat sebuah kanal baru dengan `make(chan tipe)`.
	// Kanal digolongkan oleh tipe yang dia bawa.
	pesan := make(chan string)

	// _Kirim_ sebuah nilai ke kanal dengan sintaks
	// `channel <-`. Di sini kita mengirim `"ping"`
	// ke kanal `pesan` yang kita buat di atas dari
	// gorutin baru.
	go func() { pesan <- "ping" }()

	// Sintaks `<-channel` _menerima_ sebuah nilai dari
	// kanal. Di sini kita akan menerima pesan `"ping"`
	// yang kita kirim di atas dan mencetaknya.
	psn := <-pesan
	fmt.Println(psn)
}
