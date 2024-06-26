package main

import (
	"fmt"
	"os"
	"strconv"
)

// membuat function untuk menjumlahkan total nilai di dalam array array
// dengan logic sendiri
func Sum(listData []int) int {
	var result int = 0
	for _, data := range listData {
		result = result + data
	}

	return result
}

// membuat function remove index yang sudah dijumlahkan dengan posibility
// mengambil di internet how to slice array
// sumber https://www.geeksforgeeks.org/delete-elements-in-a-slice-in-golang/
// tidak jadi menggunakan konsep dari sumber, jadinya menggunakan logic sendiri
func Remove(listData []int, removeIndex int) []int {
	result := []int{}
	for index, data := range listData {
		if index == removeIndex {
			continue
		} else {
			result = append(result, data)
		}
	}

	return result
}

// mencari index posibility terakhir di dalam array list data
// menggunakan logic sendiri

func FindIndexPosibility(listData []int, posibility int) int {
	var result int
	for i := len(listData) - 1; i > 0; i-- {
		if listData[i] < posibility {
			result = i
			break
		}
	}

	return result
}

func main() {

	// membuat function handler panic
	// referensi https://www.digitalocean.com/community/tutorials/handling-panics-in-go
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// mengambil dari project yang sudah dikerjakan sebelumnya contoh: go run main.go env-dev
	// sumber dari https://github.com/arifhdyt72/absensi-backend
	os.Args = []string{"main.go", "21", "4"}

	// konversi string to int
	// menggunakan logic sendiri
	maxTarget, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("failed to convert int value", err.Error())
		return
	}
	posibility, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("failed to convert int value", err.Error())
		return
	}

	// melakukan inisiasi data yang dibutuhkan untuk printout hasil
	listData := []int{}
	index := maxTarget - 1
	var output int = 0
	for i := 0; i < maxTarget; i++ {
		listData = append(listData, 1)
	}

	// melakukan looping data maximal target
	// menggunakan logic sendiri
	for i := 0; i < maxTarget; i++ {
		// jika looping pertama tidak ada penjumlahan dari posibility
		if i != 0 {
			// check apakah index posibility ada di index pertama
			// jika index posibility ada di index pertama maka penjumlahan posibility selesai
			if index != 0 {
				// pengecekan apakah nilai data di index tersebut < nilai posibility
				// jika iya maka lanjutkan penjumlahan posibility (data[index] + (data[index -1])
				// jika tidak lanjutkan penjumlahan posibility ke index sebelumnya
				// jika penjumlahan posibility selesai,hapus data index yang sudah dijumlahkan dengan index[posibility]
				// jika index[posibility] masih < nilai posibility, nilai index[posibility] masih sama
				// jika tidak nilai index[posibility] beralih ke index sebelumnya
				if listData[index] < posibility {
					listData[index] = listData[index] + listData[index-1]
					listData = Remove(listData, index-1)
					index = FindIndexPosibility(listData, posibility)
				} else {
					index = index - 1
					listData[index] = listData[index] + listData[index-1]
					listData = Remove(listData, index-1)
					index = FindIndexPosibility(listData, posibility)
				}
			}
		}

		// melakukan print log dari hasil penjumlahan posibility
		// konversi data array int to string
		// jumlahkan data array dengan function sum
		// print log
		// jumlahkan total print log + print log
		var printData string = ""
		for index, logData := range listData {
			s := strconv.Itoa(logData)
			if index == 0 {
				printData = printData + s
			} else {
				printData = printData + "," + s
			}
		}

		logPrint := fmt.Sprintf("%s = %d", printData, Sum(listData))
		fmt.Println(logPrint)
		output = output + 1

		// jika index[posibility] sudah di index ke 0
		// maka penjumlahan data array dengan posibility selesai
		if index == 0 {
			break
		}
	}

	// print out total log data
	fmt.Println("Output :", output)
}
