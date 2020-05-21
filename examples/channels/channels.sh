# Saat kita menjalankan program pesan `"ping"` sukses
# terkirim dari satu gorutin ke gorutin lainnya melalui
# kanal kita.
$ go run channels.go 
ping

# Normalnya pengiriman dan penerimaan akan terhambat
# sampai kedua pengirim dan penerima siap. Sifat ini
# memungkinkan kita untuk menunggu pada akhir program
# untuk pesan `"ping"` tanpa harus menggunakan sistem
# sinkronisasi lain.
