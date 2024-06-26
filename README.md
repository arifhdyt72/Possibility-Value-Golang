# Possibility-Value-Golang
Possibility Value Golang adalah mini sistem untuk melakukan perhitungan data array berdasarkan nilai maksimal target dan nilai possibility

untuk menjalannya ketik perintah di terminial   
```
go run main.go
```
di dalam mini sistem ini juga terdapat beberapa unit test yang sudah dibuat yaitu unit test untuk function yang
sudah dibuat sebelumnya dan unit test untuk 3 test nilai possibility dengan kondisi yang berbeda.

pembuatan unit test ini menggunakan library
```
https://github.com/stretchr/testify
```

pembuatan unit test ini juga menggunakan beberapa referensi dari internet dan referensi dari project yang sudah
dibuat sebelumnya. Adapun referensi yang digunakan diantaranya:
```
- https://dev.to/boncheff/table-driven-unit-tests-in-go-407b (menerapkan table test di unit test)
- https://github.com/arifhdyt72/golang_restfull_api (unit test dari project yang sudah dikerjakan sebelumnya)
```

adapun unit test yang digunakan diantaranya:
```
- Unit Test pengecekan function SUM
- Unit Test pengecekan function Remove Index
- Unit Test pengecekan function FindIndexPosibility
- Unit Test pengecekan perhitungan data array berdasarkan nilai maksimal target dan nilai possibility
```

perintah untuk menjalankan unit test
```
go test -v (untuk menjalankan semua unit test)
go test -v -run=namaFunction (untuk menjalankan unit test di function tertentu)
```