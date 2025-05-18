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

func tambahWorkout(data *tabWorkout, n *int) {
	var jenis, jadwal string
	var durasi, kalori int

	fmt.Print("Masukkan jenis olahraga: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&durasi)
	fmt.Print("Masukkan jumlah kalori: ")
	fmt.Scan(&kalori)
	fmt.Print("Masukkan jadwal (YYYY-MM-DD): ")
	fmt.Scan(&jadwal)

	w := Workout{jenis, durasi, kalori, jadwal}
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

	var jenis, jadwal string
	var durasi, kalori int

	fmt.Print("Jenis olahraga baru: ")
	fmt.Scan(&jenis)
	fmt.Print("Durasi baru (menit): ")
	fmt.Scan(&durasi)
	fmt.Print("Kalori baru: ")
	fmt.Scan(&kalori)
	fmt.Print("Jadwal baru (YYYY-MM-DD): ")
	fmt.Scan(&jadwal)

	w := Workout{jenis, durasi, kalori, jadwal}
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
	found := false
	for i := 0; i < n; i++ {
		if data[i].JenisOlahraga == keyword {
			w := data[i]
			fmt.Printf("[%d] %s | %d menit | %d kalori | %s\n", i, w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
	return found
}

func tampilkanSemuaWorkout(data tabWorkout, n int) {
	for i := 0; i < n; i++ {
		w := data[i]
		fmt.Printf("[%d] %s | %d menit | %d kalori | %s\n", i, w.JenisOlahraga, w.DurasiMenit, w.Kalori, w.Jadwal)
	}
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
