# Saat kita menjalankan program ini, pertama kita 
# melihat keluaran dari pemanggilan fungsi blocking, lalu
# penyisipan keluaran dari dua gorutin. Penyisipan ini
# memperlihatkan gorutin-gorutin sedang dijalankan secara
# bersamaan oleh runtime Go.
$ go run goroutines.go
langsung : 0
langsung : 1
langsung : 2
gorutin : 0
anonim
gorutin : 1
gorutin : 2
selesai

# Selanjutnya kita akan melihat pelengkap dari gorutin
# pada program konkuren Go: kanal (channel).
