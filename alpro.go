package main

import "fmt"

type Workout struct {
	JenisOlahraga string
	DurasiMenit   int
	Kalori        int
	Jadwal        string 
}
const NMAX int = 100
type tabWorkout [NMAX]Workout

func main() {
	var data tabWorkout
	var n, pilihan int

	for {
		pilihan = menu()

		if pilihan == 1 {
			tambahWorkout(&data, &n)
		} else if pilihan == 2 {
			editWorkout(&data, n)
		} else if pilihan == 3 {
			hapusWorkout(&data, &n)
		} else if pilihan == 4 {
			var keyword string
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scan(&keyword)
			sequentialSearch(data, n, keyword)
		} else if pilihan == 5 {
			var keyword string
			fmt.Print("Masukkan jenis olahraga (tepat): ")
			fmt.Scan(&keyword)
			binarySearch(data, n, keyword)
		} else if pilihan == 6 {
			r := rekomendasiLatihan(data, n)
			fmt.Println("\nRekomendasi Latihan:", r)
		} else if pilihan == 7 {
			statistikWorkout(data, n)
		} else if pilihan == 8 {
			sortingWorkout(&data, n)
		} else if pilihan == 9 {
			filterWorkout(data, n)
		} else if pilihan == 10 {
			tampilkanSemuaWorkout(data, n)
		} else if pilihan == 11 {
			fmt.Println("\nKeluar dari program.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}



func menu() int {
	fmt.Println("\nMenu:")
	fmt.Println("1. Tambah Workout")
	fmt.Println("2. Edit Workout")
	fmt.Println("3. Hapus Workout")
	fmt.Println("4. Cari Workout (Sequential Search)")
	fmt.Println("5. Cari Workout (Binary Search)")
	fmt.Println("6. Rekomendasi Latihan")
	fmt.Println("7. Statistik Workout")
	fmt.Println("8. Sorting Workout")
	fmt.Println("9. Filter Workout")
	fmt.Println("10. Tampilkan Semua Workout")
	fmt.Println("11. Keluar\n")
	fmt.Print("Pilih opsi: ")

	var pilihan int
	fmt.Scan(&pilihan)
	return pilihan
}

func tambahWorkout(data *tabWorkout, n *int) {
	var w Workout

	fmt.Print("\nMasukkan jenis olahraga: ")
	fmt.Scan(&w.JenisOlahraga)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&w.DurasiMenit)
	fmt.Print("Masukkan jumlah kalori: ")
	fmt.Scan(&w.Kalori)
	fmt.Print("Masukkan jadwal (YYYY-MM-DD): ")
	fmt.Scan(&w.Jadwal)

	if validasiWorkout(w) {
		data[*n] = w
		*n++
		fmt.Println("Workout berhasil ditambahkan!")
	} else {
		fmt.Println("Data tidak valid.")
	}
}

func editWorkout(data *tabWorkout, n int) {
	tampilkanSemuaWorkout(*data, n)

	var index int
	fmt.Print("Masukkan indeks workout yang ingin diedit: ")
	fmt.Scan(&index)

	if index < 0 || index >= n {
		fmt.Println("Indeks tidak valid.")
		return
	}

	var w Workout
	fmt.Print("Jenis olahraga baru: ")
	fmt.Scan(&w.JenisOlahraga)
	fmt.Print("Durasi baru (menit): ")
	fmt.Scan(&w.DurasiMenit)
	fmt.Print("Kalori baru: ")
	fmt.Scan(&w.Kalori)
	fmt.Print("Jadwal baru (YYYY-MM-DD): ")
	fmt.Scan(&w.Jadwal)

	if validasiWorkout(w) {
		data[index] = w
		fmt.Println("Data berhasil diperbarui.")
	} else {
		fmt.Println("Data tidak valid.")
	}
}

func hapusWorkout(data *tabWorkout, n *int) {
	tampilkanSemuaWorkout(*data, *n)
	var index int
	fmt.Print("Masukkan indeks workout yang ingin dihapus: ")
	fmt.Scan(&index)
	if index < 0 || index >= *n {
		fmt.Println("Indeks tidak valid.")
		return
	}
	for i := index; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Workout berhasil dihapus.")
}

func sequentialSearch(data tabWorkout, n int, keyword string) bool {
	var i int
	var found bool = false
	var w Workout

	for i = 0; i < n; i++ {
		if data[i].JenisOlahraga == keyword {
			w = data[i]
			fmt.Printf("[%d] %s | %d menit | %d kalori | %s\n", i, w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
	return found
}

func insertionSortByJenis(data *tabWorkout, n int) {
	var i, j int
	var key Workout

	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].JenisOlahraga > key.JenisOlahraga {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func binarySearch(data tabWorkout, n int, keyword string) bool {
	insertionSortByJenis(&data, n)

	var low, high, mid int
	var w Workout
	var found bool = false

	low = 0
	high = n - 1

	for low <= high && !found {
		mid = (low + high) / 2
		if data[mid].JenisOlahraga == keyword {
			w = data[mid]
			fmt.Printf("[%d] %s | %d menit | %d kalori | %s\n", mid, w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
			found = true
		} else if data[mid].JenisOlahraga < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}

	return found
}


func statistikWorkout(data tabWorkout, n int) {
	var totalKalori, start, i int
	var w Workout

	totalKalori = 0
	fmt.Println("10 Aktivitas Terakhir:")

	start = 0
	if n > 10 {
		start = n - 10
	}

	for i = start; i < n; i++ {
		w = data[i]
		fmt.Printf("%s | %d menit | %d kalori | %s\n", w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
		totalKalori = totalKalori + w.Kalori
	}

	fmt.Println("Total Kalori Terbakar:", totalKalori)
}

func insertionSortByKalori(data *tabWorkout, n int) {
	var i, j int
	var key Workout

	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].Kalori < key.Kalori {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func sortingWorkout(data *tabWorkout, n int) {
	insertionSortByKalori(data, n)
	fmt.Println("Workout diurutkan berdasarkan kalori terbanyak:")
	tampilkanSemuaWorkout(*data, n)
}

func filterWorkout(data tabWorkout, n int) {
	var keyword string
	fmt.Print("Masukkan jenis olahraga untuk filter: ")
	fmt.Scan(&keyword)
	sequentialSearch(data, n, keyword)
}

func cariIndeksJenis(jenis string, jenisUnik [NMAX]string, jumlahJenis int) int{
	var i int

	for i = 0; i < jumlahJenis; i++ {
		if jenisUnik[i] == jenis {
			return i
		}
	}
	return -1
}

func rekomendasiLatihan(data tabWorkout, n int) string{
	var jenisUnik [NMAX]string
	var frekuensi [NMAX]int
	var kaloriTertinggi [NMAX]int
	var jumlahJenis, maxIndex, i int
	var jenis string
	var kal int
	var idx int

	for i = 0; i < n; i++ {
		jenis = data[i].JenisOlahraga
		kal = data[i].Kalori
		idx = cariIndeksJenis(jenis, jenisUnik, jumlahJenis)

		if idx != -1 {
			frekuensi[idx]++
			if kal > kaloriTertinggi[idx] {
				kaloriTertinggi[idx] = kal
			}
		} else {
			jenisUnik[jumlahJenis] = jenis
			frekuensi[jumlahJenis] = 1
			kaloriTertinggi[jumlahJenis] = kal
			jumlahJenis++
		}
	}

	if jumlahJenis == 0{
		return "Tidak ada data workout!"
	}

	maxIndex = 0
	for i = 1; i < jumlahJenis; i++ {
		if frekuensi[i] > frekuensi[maxIndex] {
			maxIndex = i
		}
	}

	return jenisUnik[maxIndex]
}

func validasiWorkout(w Workout) bool {
	if w.DurasiMenit <= 0 || w.Kalori <= 0 || w.JenisOlahraga == "" || w.Jadwal == "" {
		return false
	}
	if len(w.Jadwal) != 10 || w.Jadwal[4] != '-' || w.Jadwal[7] != '-' {
		return false
	}
	return true
}

func tampilkanSemuaWorkout(data tabWorkout, n int) {
	var i int
	var w Workout

	for i = 0; i < n; i++ {
		w = data[i]
		fmt.Printf("[%d] %s | %d menit | %d kalori | %s\n", i, w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
	}
}
