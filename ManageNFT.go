/*
===============================================================================
						LAPORAN KOMENTAR PROGRAM
		Tugas Besar Aplikasi Manajemen Portofolio NFT dan Karya Digital
===============================================================================

1. Nomor dan Judul Topik:
   Topik 29	 - Aplikasi Manajemen Portofolio NFT dan Karya Digital

2. Deskripsi Masalah:
  	 Banyak seniman digital dan kolektor NFT kesulitan dalam mencatat, mengatur, dan memantau koleksi NFT mereka secara praktis.
	 Tanpa sistem yang terstruktur, mereka sulit mengetahui perubahan harga, nilai total investasi, dan riwayat transaksi.
	 Oleh karena itu, dibutuhkan aplikasi sederhana yang dapat membantu pengguna untuk mengelola portofolio NFT mereka secara efektif,
	 termasuk fitur pencarian, pengurutan, dan pelacakan data NFT.

3. Identitas Tim:
   - Mikael Eureka Anakotta - 103012400420
   - Dillah Emeylia Putri - 103012430061

4. Penjabaran Proses Pengerjaan:

    A. Fungsi yang Telah Diselesaikan:
      - Pembuatan tampilan menu utama dan sistem navigasi berbasis pilihan angka.
      - Implementasi fitur CRUD (Create, Read, Update, Delete) pada data NFT:
         > Pengguna dapat menambahkan NFT dengan data: nama, ID unik, harga, dan deskripsi.
         > Sistem ID bersifat auto increment untuk memastikan ID unik dan menghindari duplikasi manual.
         > Edit NFT dengan mengubah data berdasarkan ID.
         > Hapus NFT secara permanen dari array dengan menyesuaikan elemen array agar tetap rapat.
         > Tampilkan seluruh NFT yang tersedia di portofolio dalam format tabel.
      - Sistem pencarian NFT:
         > Mencari NFT berdasarkan ID.
         > Mencari NFT berdasarkan nama.
         > Mencari NFT dalam kisaran harga tertentu (rentang minimum dan maksimum).
         > Mencetak hasil pencarian dengan format tabel yang rapi.
      - Fitur pengurutan (sorting):
         > Mengurutkan NFT berdasarkan ID (ascending).
         > Mengurutkan NFT berdasarkan harga (ascending dan descending).
         > Dilengkapi logika pertukaran data dalam array saat proses sorting untuk menjaga konsistensi urutan.
      - Menampilkan total nilai portofolio NFT:
         > Menjumlahkan seluruh harga NFT dalam portofolio.
         > Menampilkan hasil akhir sebagai estimasi nilai kekayaan digital pengguna.
      - Histori aktivitas:
         > Menyimpan log aktivitas seperti penambahan, penghapusan, dan pengeditan NFT.
         > Menampilkan daftar histori aktivitas sesuai urutan waktu.
         > Menyediakan fitur untuk menghapus seluruh histori yang tersimpan jika pengguna menginginkannya (clear history).

    B. Permasalahan Teknis yang Dihadapi Selama Pengembangan:
      - Awalnya mengalami kesulitan saat menentukan bagaimana menyimpan banyak data NFT menggunakan array,
        karena array bersifat statis dan tidak fleksibel seperti slice.
        Solusinya adalah membuat batas maksimal array dan menggunakan variabel penghitung untuk menandai jumlah data yang aktif.

      - Menghapus elemen array membutuhkan penyesuaian agar tidak ada "lubang" kosong di tengah array.
        Maka digunakan logika pergeseran data, yaitu semua elemen setelah yang dihapus digeser satu ke kiri.

      - Input dari pengguna terkadang menyebabkan error atau panic karena tidak divalidasi dengan baik (misalnya input huruf saat diminta angka).
        Solusi sementara dilakukan dengan menambahkan pengecekan sederhana dan pengulangan input jika error terjadi.

      - Karena tidak menggunakan fitur seperti slice atau map, maka penyimpanan log aktivitas dilakukan secara manual menggunakan array string dan penghitung indeks.
        Ini cukup terbatas namun berhasil mencatat aktivitas penting seperti tambah/edit/hapus.

      - Ketika melakukan edit dan sorting secara bersamaan, perlu hati-hati agar data tidak berubah secara tidak sengaja.
        Diperlukan penataan alur agar proses edit selalu dilakukan sebelum sorting atau disimpan dalam salinan sementara.

      - Untuk menghindari ID yang sama atau bentrok, digunakan sistem ID auto increment.
        Ini memastikan setiap NFT memiliki ID unik tanpa perlu pengecekan manual.

      - Antarmuka pengguna (tampilan di terminal) masih sangat sederhana dan kurang menarik secara visual.
        Ke depannya, agar program lebih menarik di mata user, bisa dilakukan:
         > Penambahan border atau garis horizontal agar data lebih rapi.
         > Pewarnaan teks (jika terminal mendukung) untuk membedakan menu, data, dan peringatan.
         > Penggunaan spasi yang konsisten agar semua kolom sejajar.
         > Tampilan loading singkat atau animasi karakter agar lebih hidup (dengan teknik sederhana).

	C. Rencana Pengembangan Selanjutnya:
      - Penambahan validasi input yang lebih kuat:
         > Mendeteksi karakter yang tidak valid pada nama NFT.
         > Membatasi ID agar tidak boleh sama.
         > Menolak input kosong atau salah format.
      - Penyempurnaan tampilan data agar lebih interaktif dan mudah dibaca.
      - Menambahkan fitur ekspor data ke file teks.
      - Implementasi pencarian dan pengurutan gabungan (multi-parameter).
===============================================================================
*/

package main

import "fmt"

const max = 1204

type NFT struct {
	nama         string
	id           int
	HargaAwalNFT float64
	harga        float64
}

var dataNFT arrNFT
var jumlahNFT int = 0
var history historyArr
var IDmasuk int = 100
var nHistory int = 0

type arrNFT [max]NFT
type historyArr [max]string

//fungsi utama yang menjalankan aplikasi dan menampilkan menu utama secara berulang hingga user memilih keluar.
func main() {
	Welcome()
	for {
		pilihan := menuUtama()
		switch pilihan {
		case 1:
			tambahNFT(&dataNFT, &jumlahNFT, &IDmasuk, &history, &nHistory)
		case 2:
			editNFT(dataNFT, jumlahNFT, &history, &nHistory)
		case 3:
			hapusNFT(&dataNFT, &jumlahNFT, &history, &nHistory)
		case 4:
			menuCariNFT(dataNFT, jumlahNFT)
		case 5:
			menuUrutNFT(&dataNFT, jumlahNFT)
		case 6:
			tampilkanHistory(history, nHistory)
			tampilkanTotal(dataNFT, jumlahNFT)
		case 7:
			hapusHistory(&history, &nHistory)
		case 8:
			cetakSemuaData(dataNFT, jumlahNFT)
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Welcome menampilkan pesan sambutan saat aplikasi dijalankan.
func Welcome() {
	fmt.Println()
	fmt.Println("+==================================================+")
	fmt.Printf("|%-17s%-33s|", " ", "~SELAMAT DATANG~")
	fmt.Printf("\n|%-8s%-42s|\n", " ", "Aplikasi Manajemen NFT Portofolio!")
	fmt.Println("+==================================================+")
}

// menuUtama menampilkan menu utama dan menerima input pilihan user.
func menuUtama() int {
	fmt.Println("\n+==================== M E N U =====================+")
	fmt.Printf("|%-50s|", "1. Tambah NFT")
	fmt.Printf("\n|%-50s|", "2. Edit NFT")
	fmt.Printf("\n|%-50s|", "3. Hapus NFT")
	fmt.Printf("\n|%-50s|", "4. Cari NFT")
	fmt.Printf("\n|%-50s|", "5. Urutkan NFT")
	fmt.Printf("\n|%-50s|", "6. Tampilkan History dan Total Nilai Portofolio")
	fmt.Printf("\n|%-50s|", "7. Hapus History")
	fmt.Printf("\n|%-50s|", "8. Cetak Semua Data")
	fmt.Printf("\n|%-50s|", "9. Keluar")
	fmt.Printf("\n+==================================================+\n")

	var pilihan int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Println()
	return pilihan
}

// tambahNFT digunakan untuk menambah data NFT baru ke dalam portofolio. 
// Fungsi ini meminta input nama dan harga NFT dari user, mengatur ID secara auto increment, 
// serta mencatat aktivitas penambahan ke dalam history.
func tambahNFT(dataNFT *arrNFT, jumlahNFT *int, IDmasuk *int, history *historyArr, nHistory *int) {
	if *jumlahNFT >= max {
		fmt.Println("Data penuh, tidak bisa menambahkan NFT lagi.")
		return
	}

	fmt.Println("+==============================================+")
	fmt.Println("|           Tambah NFT ke Portofolio           |")
	fmt.Println("+==============================================+")
	*IDmasuk++
	dataNFT[*jumlahNFT].id = *IDmasuk
	fmt.Printf("| ID NFT   : %-33d |\n", dataNFT[*jumlahNFT].id)
	fmt.Println("+==============================================+")
	fmt.Printf("| Nama NFT : ")
	fmt.Scan(&dataNFT[*jumlahNFT].nama)
	fmt.Printf("| Harga NFT: ")
	fmt.Scan(&dataNFT[*jumlahNFT].harga)
	fmt.Println("+==============================================+")

	history[*nHistory] = fmt.Sprintf("Menambahkan NFT ID %d", *IDmasuk)
	*nHistory += 1
	*jumlahNFT++
}

// editNFT digunakan untuk mengedit data NFT yang sudah ada berdasarkan ID. 
// User dapat mengubah nama, harga, atau ID NFT. Setiap perubahan dicatat ke dalam history.
func editNFT(dataNFT arrNFT, jumlahNFT int, history *historyArr, nHistory *int) {
	var id int
	fmt.Print("Masukkan ID NFT yang ingin diedit: ")
	fmt.Scan(&id)

	for i := 0; i <= jumlahNFT; i++ {
		if dataNFT[i].id == id {
			var pilihan int
			var newNama string
			var newHarga float64

			fmt.Println("Data NFT Ditemukan:")
			fmt.Printf("Nama: %s | Harga: %.2f\n", dataNFT[i].nama, dataNFT[i].harga)
			fmt.Println("Apa yang ingin Anda ubah?")
			fmt.Println("1. Nama NFT")
			fmt.Println("2. Harga NFT")
			fmt.Println("3. ID NFT")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&newNama)
				fmt.Println("\nPerubahan:")
				fmt.Printf("Sebelum: %s\n", dataNFT[i].nama)
				fmt.Printf("Sesudah: %s\n", newNama)
			} else if pilihan == 2 {
				fmt.Print("Masukkan harga baru: ")
				fmt.Scan(&newHarga)
				fmt.Println("\nPerubahan:")
				fmt.Printf("Sebelum: %.2f\n", dataNFT[i].harga)
				fmt.Printf("Sesudah: %.2f\n", newHarga)
			} else if pilihan == 3 {
				editIDNFT(&dataNFT, jumlahNFT+1, history, nHistory)
				return
			} else {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			var konfirmasi string
			fmt.Print("Terima perubahan? (ya/tidak): ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == "ya" {
				if pilihan == 1 {
					dataNFT[i].nama = newNama
				} else if pilihan == 2 {
					dataNFT[i].harga = newHarga
				}
				history[*nHistory] = fmt.Sprintf("Mengedit NFT ID %d", id)
				*nHistory += 1
				fmt.Println("Perubahan berhasil disimpan.")
			} else {
				fmt.Println("Perubahan dibatalkan.")
			}
			return
		}
	}
	fmt.Println("NFT tidak ditemukan.")
}

// hapusNFT digunakan untuk menghapus data NFT berdasarkan ID. 
// Setelah NFT dihapus, data di array akan digeser agar tetap rapat, dan aktivitas di-log ke history.
func hapusNFT(dataNFT *arrNFT, jumlahNFT *int, history *historyArr, nHistory *int) {
	var id int
	fmt.Print("Masukkan ID NFT yang ingin dihapus: ")
	fmt.Scan(&id)

	for i := 0; i < *jumlahNFT; i++ {
		if dataNFT[i].id == id {
			for j := i; j < *jumlahNFT-1; j++ {
				dataNFT[j] = dataNFT[j+1]
			}
			*jumlahNFT--
			history[*nHistory] = "Menghapus NFT"
			*nHistory++
			fmt.Println("NFT berhasil dihapus.")
			cetakSemuaData(*dataNFT, *jumlahNFT)
			return
		}
	}
	fmt.Println("NFT tidak ditemukan.")
}

// tampilkanHistory menampilkan seluruh riwayat aktivitas yang telah dilakukan user pada aplikasi.
func tampilkanHistory(history historyArr, nHistory int) {
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

// tampilkanTotal menghitung dan menampilkan total nilai portofolio NFT yang dimiliki user.
func tampilkanTotal(dataNFT arrNFT, jumlahNFT int) {
	var total float64
	for i := 0; i < jumlahNFT; i++ {
		total += dataNFT[i].harga
	}
	fmt.Printf("Total nilai portofolio NFT: Rp,%.2f\n", total)
	fmt.Println()
}

// menuCariNFT menampilkan menu pencarian NFT dan memanggil fungsi pencarian sesuai pilihan user.
func menuCariNFT(dataNFT arrNFT, jumlahNFT int) {
	var pilihan int
	fmt.Printf("+=================== PENCARIAN ====================+\n")
	fmt.Printf("|%-52s|\n", "   Cari berdasarkan : ")
	fmt.Printf("|%-52s|\n", "1. Sequential Search (Nama)")
	fmt.Printf("|%-52s|\n", "2. Binary Search (ID)")
	fmt.Printf("|%-52s|\n", "3. Sequential Search (Harga)")
	fmt.Printf("|%-52s|\n", "4. Filter Harga dalam Rentang")
	fmt.Printf("+==================================================+\n")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Masukkan nama NFT: ")
		fmt.Scan(&nama)
		sequentialSearch(dataNFT, jumlahNFT, nama)
	case 2:
		var id int
		fmt.Print("Masukkan ID NFT: ")
		fmt.Scan(&id)
		binarySearch(dataNFT, jumlahNFT, id)
	case 3:
		var harga float64
		fmt.Print("Masukkan harga NFT: ")
		fmt.Scan(&harga)
		sequentialSearchHarga(dataNFT, jumlahNFT, harga)
	case 4:
		var min, max float64
		fmt.Print("Masukkan harga minimum: ")
		fmt.Scan(&min)
		fmt.Print("Masukkan harga maksimum: ")
		fmt.Scan(&max)
		filterNFT(dataNFT, jumlahNFT, min, max)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	fmt.Println()
}

// sequentialSearch mencari NFT berdasarkan nama menggunakan metode pencarian linear.
func sequentialSearch(dataNFT arrNFT, jumlahNFT int, nama string) {
	found := false
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].nama == nama {
			fmt.Printf("|%-14s %-35d|\n", "ID NFT: ", dataNFT[i].id)
			fmt.Printf("|%-14s %-35s|\n", "Nama NFT: ", dataNFT[i].nama)
			fmt.Printf("|%-14s %-35.2f|\n", "Harga NFT: ", dataNFT[i].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("NFT tidak ditemukan.")
	}
	fmt.Println()
}

// sequentialSearchHarga mencari NFT berdasarkan harga menggunakan metode pencarian linear.
func sequentialSearchHarga(dataNFT arrNFT, jumlahNFT int, harga float64) {
	found := false
	for i := 0; i < jumlahNFT && !found; i++ {
		if dataNFT[i].harga == harga {
			fmt.Printf("Ditemukan: ID: %d, Nama: %s, Harga: %.2f\n", dataNFT[i].id, dataNFT[i].nama, dataNFT[i].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("NFT dengan harga tersebut tidak ditemukan.")
	}
}

// binarySearch mencari NFT berdasarkan ID menggunakan metode binary search. 
// Data diurutkan terlebih dahulu berdasarkan ID sebelum pencarian.
func binarySearch(dataNFT arrNFT, jumlahNFT int, cari int) {
	selectionSortByID(&dataNFT, jumlahNFT)
	kiri := 0
	kanan := jumlahNFT - 1
	found := false
	idx := -1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if dataNFT[mid].id == cari {
			found = true
			idx = mid
			kiri = kanan + 1
		} else if dataNFT[mid].id < cari {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	if found {
		fmt.Printf("Ditemukan: ID: %d, Nama: %s, Harga: %.2f\n", dataNFT[idx].id, dataNFT[idx].nama, dataNFT[idx].harga)
	} else {
		fmt.Println("NFT tidak ditemukan.")
	}
	fmt.Println()
}

// menuUrutNFT menampilkan menu pengurutan NFT dan memanggil fungsi sorting sesuai pilihan user.
func menuUrutNFT(dataNFT *arrNFT, jumlahNFT int) {
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
	case 1:
		switch pilihanMetode {
		case 1:
			selectionSortByID(dataNFT, jumlahNFT)
		case 2:
			insertionSortByID(dataNFT, jumlahNFT)
		default:
			fmt.Println("Metode tidak valid.")
			return
		}
	case 2:
		switch pilihanMetode {
		case 1:
			selectionSortByHarga(dataNFT, jumlahNFT)
		case 2:
			insertionSortByHarga(dataNFT, jumlahNFT)
		default:
			fmt.Println("Metode tidak valid.")
			return
		}
	default:
		fmt.Println("Pilihan atribut tidak valid.")
		return
	}

	fmt.Printf("+======+=========================+==============+\n")
	fmt.Printf("| %-4s | %-23s | %-12s |\n", " ID", " Nama", " Harga")
	fmt.Printf("+======+=========================+==============+\n")
	for i := 0; i < jumlahNFT; i++ {
		fmt.Printf("| %-4d | %-23s | Rp,%-6.2f |\n", dataNFT[i].id, dataNFT[i].nama, dataNFT[i].harga)
		fmt.Printf("+======+=========================+==============+\n")
	}
	fmt.Println()
}

// selectionSortByID mengurutkan data NFT berdasarkan ID secara ascending menggunakan selection sort.
func selectionSortByID(dataNFT *arrNFT, jumlahNFT int) {
	var temp NFT
	for i := 0; i < jumlahNFT-1; i++ {
		min := i
		for j := i + 1; j < jumlahNFT; j++ {
			if dataNFT[j].id < dataNFT[min].id {
				min = j
			}
		}
		temp = dataNFT[i]
		dataNFT[i] = dataNFT[min]
		dataNFT[min] = temp
	}
	fmt.Println()
}

// insertionSortByID mengurutkan data NFT berdasarkan ID secara ascending menggunakan insertion sort.
func insertionSortByID(dataNFT *arrNFT, jumlahNFT int) {
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

// selectionSortByHarga mengurutkan data NFT berdasarkan harga secara ascending menggunakan selection sort.
func selectionSortByHarga(dataNFT *arrNFT, jumlahNFT int) {
	var temp NFT
	for i := 0; i < jumlahNFT-1; i++ {
		min := i
		for j := i + 1; j < jumlahNFT; j++ {
			if dataNFT[j].harga < dataNFT[min].harga {
				min = j
			}
		}
		temp = dataNFT[i]
		dataNFT[i] = dataNFT[min]
		dataNFT[min] = temp
	}
	fmt.Println()
}

// insertionSortByHarga mengurutkan data NFT berdasarkan harga secara ascending menggunakan insertion sort.
func insertionSortByHarga(dataNFT *arrNFT, jumlahNFT int) {
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

// editIDNFT digunakan untuk mengubah ID NFT tertentu, dengan validasi agar ID baru tidak sama dengan ID lain.
func editIDNFT(dataNFT *arrNFT, jumlahNFT int, history *historyArr, nHistory *int) {
	var oldID, newID int
	fmt.Print("Masukkan ID NFT yang ingin diubah: ")
	fmt.Scan(&oldID)

	idx := -1
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].id == oldID {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("NFT dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan ID baru: ")
	fmt.Scan(&newID)

	// Cek apakah ID baru sudah digunakan
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].id == newID {
			fmt.Println("ID baru sudah digunakan oleh NFT lain.")
			return
		}
	}

	old := dataNFT[idx].id
	dataNFT[idx].id = newID
	history[*nHistory] = fmt.Sprintf("Mengubah ID NFT dari %d ke %d", old, newID)
	*nHistory++
	fmt.Println("ID NFT berhasil diubah.")
}

// cetakSemuaData menampilkan seluruh data NFT yang ada di portofolio dalam format tabel.
func cetakSemuaData(dataNFT arrNFT, jumlahNFT int) {
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

// hapusHistory menampilkan menu penghapusan history dan memanggil fungsi sesuai pilihan user.
func hapusHistory(history *historyArr, nHistory *int) {
	var pilihan int
	fmt.Println("Pilihan penghapusan history:")
	fmt.Println("1. Hapus history tertentu")
	fmt.Println("2. Hapus semua history")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		hapusHistoryTertentu(history, nHistory)
	case 2:
		hapusSemuaHistory(nHistory)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// hapusHistoryTertentu menghapus salah satu history berdasarkan nomor yang dipilih user.
func hapusHistoryTertentu(history *historyArr, nHistory *int) {
	if *nHistory == 0 {
		fmt.Println("Tidak ada history untuk dihapus")
		return
	}

	for i := 0; i < *nHistory; i++ {
		fmt.Printf("%d. %s\n", i+1, history[i])
	}

	var nomor int
	fmt.Print("Masukkan nomor history yang ingin dihapus: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > *nHistory {
		fmt.Println("Nomor history tidak valid")
		return
	}

	for i := nomor - 1; i < *nHistory-1; i++ {
		history[i] = history[i+1]
	}
	*nHistory--
	fmt.Println("History berhasil dihapus")
}

// hapusSemuaHistory menghapus seluruh history aktivitas yang tersimpan.
func hapusSemuaHistory(nHistory *int) {
	if *nHistory == 0 {
		fmt.Println("Tidak ada history untuk dihapus")
		return
	}

	var konfirmasi string
	fmt.Print("Yakin ingin menghapus semua history? (ya/tidak): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "ya" {
		*nHistory = 0
		fmt.Println("Semua history telah dihapus")
	} else {
		fmt.Println("Penghapusan history dibatalkan")
	}
}

// filterNFT menampilkan NFT yang memiliki harga dalam rentang tertentu (min sampai max).
func filterNFT(dataNFT arrNFT, jumlahNFT int, min, max float64) {
	found := false
	fmt.Println("+============== NFT Dalam Rentang Harga ==============+")
	fmt.Printf("+======+=========================+====================+\n")
	fmt.Printf("| %-4s | %-23s | %-18s |\n", "ID", "Nama", "Harga")
	fmt.Printf("+======+=========================+====================+\n")
	for i := 0; i < jumlahNFT; i++ {
		if dataNFT[i].harga >= min && dataNFT[i].harga <= max {
			nama := dataNFT[i].nama
			if len(nama) > 23 {
				nama = nama[:20] + "..."
			}
			fmt.Printf("| %-4d | %-23s | Rp,%-15.2f |\n", dataNFT[i].id, nama, dataNFT[i].harga)
			fmt.Printf("+======+=========================+====================+\n")
			found = true
		}
	}
	if !found {
		fmt.Println("|            Tidak ada NFT dalam rentang harga tersebut             |")
		fmt.Printf("+===============================================================+\n")
	}
	fmt.Println()
}
