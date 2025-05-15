package main

import "fmt"

const max = 1000

// Struktur NFT menyimpan nama, ID, dan harga NFT
type NFT struct {
	nama  string
	id    int
	harga float64
}

type arrNFT [max]NFT
type historyArr [max]string

var dataNFT arrNFT
var jumlahNFT int = -1
var history historyArr
var nHistory int

func Welcome() {
	fmt.Println()
	fmt.Println("+==================================================+")
	fmt.Printf("|%-14s%-36s|", " ", "SELAMAT DATANG SENSEI~")
	fmt.Printf("\n|%-8s%-42s|\n", " ", "Aplikasi Manajemen NFT Portofolio!")
	fmt.Println("+==================================================+")
}

func main() {
	var pilihan int
	Welcome()
	fmt.Println()
	for pilihan != 8 {
		// Tampilkan menu utama
		fmt.Println("+==================== M E N U =====================+")
		fmt.Printf("|%-50s|", "1. Tambah NFT")
		fmt.Printf("\n|%-50s|", "2. Edit NFT")
		fmt.Printf("\n|%-50s|", "3. Hapus NFT")
		fmt.Printf("\n|%-50s|", "4. Cari NFT")
		fmt.Printf("\n|%-50s|", "5. Urutkan NFT")
		fmt.Printf("\n|%-50s|", "6. Tampilkan History dan Total Nilai Portofolio")
		fmt.Printf("\n|%-50s|", "7. Cetak Semua Data")
		fmt.Printf("\n|%-50s|", "8. Keluar")
		fmt.Printf("\n+==================================================+\n")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			tambahNFT()
		case 2:
			editNFT()
		case 3:
			hapusNFT()
		case 4:
			menuCariNFT()
		case 5:
			menuUrutNFT()
		case 6:
			tampilkanHistory()
			tampilkanTotal()
		case 7:
			cetakSemuaData()
			fmt.Println()
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			fmt.Println()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menambahkan NFT
func tambahNFT() {
	if jumlahNFT >= max {
		fmt.Println("Data penuh, tidak bisa menambahkan NFT lagi.")
		return
	}
	jumlahNFT++
	fmt.Print("ID NFT: ")
	fmt.Scan(&dataNFT[jumlahNFT].id)
	fmt.Print("Nama NFT : ")
	fmt.Scan(&dataNFT[jumlahNFT].nama)
	fmt.Print("Harga NFT: ")
	fmt.Scan(&dataNFT[jumlahNFT].harga)
	fmt.Println()
	history[nHistory] = "Menambahkan NFT"
	nHistory++
}

// Fungsi untuk mengedit data NFT berdasarkan ID
func editNFT() {
	var id int
	fmt.Print("Masukkan ID NFT yang ingin diedit: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].id == id {
			fmt.Print("Nama Baru (1 kata): ")
			fmt.Scan(&dataNFT[i].nama)
			fmt.Print("Harga Baru: ")
			fmt.Scan(&dataNFT[i].harga)
			history[nHistory] = "Mengedit NFT"
			nHistory++
			return
		}
	}
	fmt.Println("NFT tidak ditemukan.")
	fmt.Println()
}

// Fungsi untuk menghapus NFT berdasarkan ID
func hapusNFT() {
	var id int
	fmt.Print("Masukkan ID NFT yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].id == id {
			for j := i; j < jumlahNFT-1; j++ {
				dataNFT[j] = dataNFT[j+1]
			}
			jumlahNFT--
			history[nHistory] = "Menghapus NFT"
			nHistory++
			fmt.Println("NFT berhasil dihapus.")
			fmt.Println()
			cetakSemuaData()
			return
		}
	}
	fmt.Println("NFT tidak ditemukan.")
	fmt.Println()
}

// Fungsi untuk menampilkan semua aktivitas
func tampilkanHistory() {
	fmt.Println("=== Riwayat Aktivitas ===")
	if nHistory == 0 {
		fmt.Println("Belum ada aktivitas.")
	} else {
		for i := 0; i < nHistory; i++ {
			fmt.Printf("%d. %s\n", i+1, history[i])
		}
	}
	fmt.Println()
}

// Hitung total nilai semua NFT
func tampilkanTotal() {
	var total float64
	total = 0
	for i := 0; i < jumlahNFT; i++ {
		total += dataNFT[i].harga
	}
	fmt.Printf("Total nilai portofolio NFT: Rp,%.2f\n", total)
	fmt.Println()
}

// Menu pemilihan metode pencarian
func menuCariNFT() {
	var pilihan int
	fmt.Printf("+======================= PENCARIAN ========================+\n")
	fmt.Printf("|%-52s|\n", "1. Sequential Search (Nama)")
	fmt.Printf("|%-52s|\n", "2. Binary Search (ID)")
	fmt.Printf("|%-52s|\n", "3. Sequential Search (Harga)")
	fmt.Printf("+==========================================================+\n")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Masukkan nama NFT: ")
		fmt.Scan(&nama)
		sequentialSearch(nama)
	case 2:
		var id int
		fmt.Print("Masukkan ID NFT: ")
		fmt.Scan(&id)
		binarySearch(id)
	case 3:
		var harga float64
		fmt.Print("Masukkan harga NFT: ")
		fmt.Scan(&harga)
		sequentialSearchHarga(harga)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	fmt.Println()
}

// Sequential search: cari berdasarkan nama dari NFT
func sequentialSearch(nama string) {
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].nama == nama {
			fmt.Printf("|%-14s %-35d|\n", "ID NFT: ", dataNFT[i].id)
			fmt.Printf("|%-14s %-35s|\n", "Nama NFT: ", dataNFT[i].nama)
			fmt.Printf("|%-14s %-35f|\n", "Harga NFT: ", dataNFT[i].harga)
			return
		}
	}
	fmt.Println("NFT tidak ditemukan.")
	fmt.Println()
}

// Binary search: cari berdasarkan ID
func binarySearch(target int) {
	selectionSortByID()
	low := 0
	high := jumlahNFT - 1
	for low <= high {
		mid := (low + high) / 2
		if dataNFT[mid].id == target {
			fmt.Printf("Ditemukan: ID: %d, Nama: %s, Harga: %f\n", dataNFT[mid].id, dataNFT[mid].nama, dataNFT[mid].harga)
			return
		} else if dataNFT[mid].id < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("NFT tidak ditemukan.")
	fmt.Println()
}

// Menu metode sortnya
func menuUrutNFT() {
	var pilihanAtribut, pilihanMetode int

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. ID")
	fmt.Println("2. Harga")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihanAtribut)

	fmt.Println("\nPilih metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihanMetode)
	fmt.Println()

	switch pilihanAtribut {
	case 1: // By ID
		switch pilihanMetode {
		case 1:
			selectionSortByID()
		case 2:
			insertionSortByID()
		default:
			fmt.Println("Metode tidak valid.")
			return
		}
	case 2: // By Harga
		switch pilihanMetode {
		case 1:
			selectionSortByHarga()
		case 2:
			insertionSortByHarga()
		default:
			fmt.Println("Metode tidak valid.")
			return
		}
	default:
		fmt.Println("Pilihan atribut tidak valid.")
		return
	}

	// Cetak hasil
	fmt.Printf("+======+=========================+==============+\n")
	fmt.Printf("| %-4s | %-23s | %-12s |\n", " ID", " Nama", " Harga")
	fmt.Printf("+======+=========================+==============+\n")
	for i := 0; i < jumlahNFT; i++ {
		fmt.Printf("| %-4d | %-23s | Rp,%-6.2f |\n", dataNFT[i].id, dataNFT[i].nama, dataNFT[i].harga)
		fmt.Printf("+======+=========================+==============+\n")
	}
	fmt.Println()
}

func insertionSortByID() {
	for i := 1; i < jumlahNFT; i++ {
		temp := dataNFT[i]
		j := i - 1
		for j >= 0 && dataNFT[j].id > temp.id {
			dataNFT[j+1] = dataNFT[j]
			j--
		}
		dataNFT[j+1] = temp
	}
}

// Seleksi berdasarkan harga
func selectionSortByHarga() {
	for i := 0; i < jumlahNFT-1; i++ {
		min := i
		for j := i + 1; j < jumlahNFT; j++ {
			if dataNFT[j].harga < dataNFT[min].harga {
				min = j
			}
		}
		dataNFT[i], dataNFT[min] = dataNFT[min], dataNFT[i]
	}
	fmt.Println()
}

// Cari NFT berdasarkan harga (sequential search)
func sequentialSearchHarga(harga float64) {
	found := false
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].harga == harga {
			fmt.Printf("Ditemukan: ID: %d, Nama: %s, Harga: %.2f\n", dataNFT[i].id, dataNFT[i].nama, dataNFT[i].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("NFT dengan harga tersebut tidak ditemukan.")
	}
}

// Insertion sort berdasarkan harga
func insertionSortByHarga() {
	for i := 1; i < jumlahNFT; i++ {
		temp := dataNFT[i]
		j := i - 1
		for j >= 0 && dataNFT[j].harga > temp.harga {
			dataNFT[j+1] = dataNFT[j]
			j--
		}
		dataNFT[j+1] = temp
	}
	fmt.Println()
}

// Seleksi berdasarkan ID
func selectionSortByID() {
	for i := 0; i < jumlahNFT-1; i++ {
		min := i
		for j := i + 1; j < jumlahNFT; j++ {
			if dataNFT[j].id < dataNFT[min].id {
				min = j
			}
		}
		dataNFT[i], dataNFT[min] = dataNFT[min], dataNFT[i]
	}
	fmt.Println()
}

func cetakSemuaData() { //buat ngeprint semua NFT yang kita punya
	if jumlahNFT == -1 {
		fmt.Println("+==================================================+")
		fmt.Printf("|%-14s%-36s|\n", " ", "Anda belum mempunyai NFT")
		fmt.Println("+==================================================+")
		return
	}
	fmt.Println("+============== Daftar NFT Milikmu! ==============+")
	fmt.Printf("+======+=========================+==============+\n")
	fmt.Printf("| %-4s | %-23s | %-12s |\n", " ID", " Nama", " Harga")
	fmt.Printf("+======+=========================+==============+\n")
	for i := 0; i < jumlahNFT; i++ {
		fmt.Printf("| %-4d | %-23s | Rp,%-6.2f |\n", dataNFT[i].id, dataNFT[i].nama, dataNFT[i].harga)
		fmt.Printf("+======+=========================+==============+\n")
	}
	fmt.Println()
}
