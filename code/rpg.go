
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ==========================
// Global Variable dan struct
// ==========================
// struct ini akan digunakan untuk menyimpan data hasil bencmark
type DataBenchmark struct {
	Level          int
	DurasiIteratif float64
	DurasiRekursif float64
}

var (
	scanner     = bufio.NewScanner(os.Stdin) // scannner input global
	dataRiwayat []DataBenchmark              // slice array dinamis untuk simpan data history
)

// =========================
// main program
// =========================
func main() {
	clearScreen()
	fmt.Print("Inisialisasi Sistem RPG....")
	time.Sleep(1 * time.Second)

	for {
		clearScreen()
		printHeader()
		printMenu()

		// baca input user
		pilihanUser := bacaString("Pilihan Anda [1-5]: ")

		switch pilihanUser {
		case "1":
			jalankanModeAdventure()
		case "2":
			jalankanModeBenchmark()
		case "3":
			tampilkanPanduanAlgoritma()
		case "4":
			jalankanExportExcel()
		case "5":
			fmt.Println("\n Menutup Sistem RPG Level Architect..")
			fmt.Println("   Sampai jumpa, Terimakasih Telah Menggunakan Aplikasi kami.")
			time.Sleep(1 * time.Second) // jeda 1 detik sebelum keluar
			os.Exit(0)
		default:
			fmt.Println("\n❌ Pilihan tidak valid! Silakan pilih 1-5.")
			time.Sleep(1 * time.Second)
		}
	}
}

// =====================
// UI Atau tampilan menu
// =====================

func printHeader() {
	fmt.Println("==========================================================")
	fmt.Println("               ⚔️  RPG LEVEL ARCHITECT  ⚔️")
	fmt.Println("         Tribonacci-Based XP System Simulator (CLI)")
	fmt.Println("==========================================================")
}

func printMenu() {
	fmt.Println("1. 🛡️  Adventure Mode       - Simulasi Leveling Hero")
	fmt.Println("2. 📊  Benchmark Mode       - Uji Performa Algoritma")
	fmt.Println("3. 📜  Algorithm Guide      - Info Teknis Tribonacci")
	fmt.Println("4. 💾  Export to Excel      - Simpan hasil benchmark ke CSV  ")
	fmt.Println("5. ❌  Exit Program         - Keluar dari aplikasi")
	fmt.Println("==========================================================")
}

// ==========================
// Fitur Adventure Mode
// ==========================
func jalankanModeAdventure() {
	clearScreen()
	fmt.Println("--- 🛡️ ADVENTURE MODE: HERO LEVELING ---")
	fmt.Println("Simulasi: Hitung total XP untuk mencapai level N")

	// (1,50) artinya min max
	levelTarget := bacaInt("Masukkan Target Level Hero (1-50): ", 1, 50)

	//loading seperti game
	fmt.Printf("\nSistem Menghitung XP untuk Level %d...\n", levelTarget)
	fmt.Println("[Proses] ⚔️  Sedang mengumpulkan data quest...")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("[Proses] ✨  Kalkulasi Tribonacci Matrix...")
	time.Sleep(300 * time.Millisecond)

	//panggil algoritma disini
	hasilXP := tribonacciIteratif(levelTarget)

	fmt.Println("\n✅ HASIL KALKULASI:")
	fmt.Printf("Untuk menjadi Hero Level %d, dibutuhkan total:\n", levelTarget)
	fmt.Printf("<< %d XP >> \n", hasilXP)

	waitForEnter()

}

// ==========================
// Fitur Benchmark Mode (digunakan analisis perbandingan waktu)
// ===========================
func jalankanModeBenchmark() {
	clearScreen()

	fmt.Println("\n========================================")
	fmt.Println("   📊  BENCHMARK MODE: DATA COLLECTOR")
	fmt.Println("========================================")
	fmt.Println("⚠️  PERINGATAN: Input N > 35 semakin besar akan membuat rekursif lambat")

	// Input batas uji
	batasLevel := bacaInt("Masukkan Batas Uji level (min 5 max 50): ", 5, 50)

	fmt.Println("\nSedang memproses benchmark... Mohon tunggu.")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("%-10s | %-25s | %-25s\n", "LEVEL (N)", "ITERATIF (Seconds)", "REKURSIF (Seconds)")
	fmt.Println("-------------------------------------------------------------")

	// kosongkan riwayat
	dataRiwayat = []DataBenchmark{}

	// loop kelipatan
	for n := 5; n <= batasLevel; n += 5 {
		// uji iteratif
		const repeat = 10000000
		startIter := time.Now() // mulai stopwatch
		for i := 0; i < repeat; i++ {
			tribonacciIteratif(n) // jalankan fungsi
		}
		waktuIter := time.Since(startIter).Seconds() / float64(repeat) //stop

		// uji rekursif
		startRec := time.Now()
		_ = tribonacciRekursif(n)
		waktuRec := time.Since(startRec).Seconds()

		//tampilkan
		fmt.Printf("%-10d | %-25.9f | %-25.9f\n", n, waktuIter, waktuRec)

		// simpan data ke memori untuk menu export
		dataBaru := DataBenchmark{
			Level:          n,
			DurasiIteratif: waktuIter,
			DurasiRekursif: waktuRec,
		}
		dataRiwayat = append(dataRiwayat, dataBaru)
	}

	fmt.Println("Benchmark Selesai! Data tersimpan sementara di memori.")
	fmt.Println(" Coba Pilih Menu 4 untuk menyimpan data ini ke file Excel (CSV).")
	waitForEnter()

}

// ==========================
// Fitur algorithm guide
// ==========================
func tampilkanPanduanAlgoritma() {
	clearScreen()
	fmt.Println("\n=======================================")
	fmt.Println("   📜  ALGORITHM GUIDE: TRIBONACCI")
	fmt.Println("========================================")
	fmt.Println("Algoritma Tribonacci adalah pola deret bilangan di mana setiap suku baru adalah hasil penjumlahan dari tiga suku sebelumnya")
	fmt.Println("menjadikannya generalisasi dari deret Fibonacci yang hanya menjumlahkan dua suku sebelumnya")
	fmt.Println("Rumus Algoritmanya: T(n) = T(n-1) + T(n-2) + T(n-3)")
	fmt.Println("\n[Kompleksitasnya]")
	fmt.Println("1. Iteratif: O(n) -> Cepat (Linear)")
	fmt.Println("2. Rekursif: O(3^n) -> Lambat (Eksponensial)")

	fmt.Println("Note karena saya menggunakan golang, mungkin lebih cepat untuk iteratifnya dibandingkan dengan rekursif yang lama,")
	fmt.Println("ini tergantung dengan sistem, laptop, ataupun lainnya")

	waitForEnter()

}

// ============
// fitur export
// =============
func jalankanExportExcel() {
	clearScreen()
	fmt.Println("\n========================================")
	fmt.Println("   💾  EXPORT DATA TO EXCEL (CSV)")
	fmt.Println("========================================")

	// validasi apakah data ada dimemori
	if len(dataRiwayat) == 0 {
		fmt.Println("⚠️  Data kosong! Silakan jalankan 'Benchmark Mode' (Menu 2) dulu.")
		waitForEnter()
		return
	}

	fileNama := "hasil_benchmark_rpg.csv"
	file, err := os.Create(fileNama)
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		waitForEnter()
		return
	}
	defer file.Close()

	// tulis header CSV
	fmt.Fprint(file, "Level (N), waktu iteratif (Second), Waktu Rekursif (Seconds)\n")

	for _, row := range dataRiwayat {
		fmt.Fprintf(file, "%d,%.10f,%.10f\n", row.Level, row.DurasiIteratif, row.DurasiRekursif)
	}

	fmt.Printf("Berhasil diexport ke file: %s\n", fileNama)
	fmt.Println("Tips: Buka file ini di Excel, lalu Insert > Chart untuk grafik!")

	waitForEnter()
}

// ==========================
// ALGORITGN INTI TRIBONACCI
// ==========================

// Ini fungsi iteratif
func tribonacciIteratif(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1

	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

// Ini fungsi rekursif
func tribonacciRekursif(n int) int {
	// Base Case (Kondisi Berhenti)
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	// Recursive Step
	return tribonacciRekursif(n-1) + tribonacciRekursif(n-2) + tribonacciRekursif(n-3)
}

// =====================================
// Fungsi bantuan atau pendukung lainnya
// =====================================
// function yang digunakan untuk baca string
func bacaString(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// function baca integer
func bacaInt(prompt string, min, max int) int {
	for {
		inputString := bacaString(prompt)
		val, err := strconv.Atoi(inputString)

		if err != nil {
			// jika error atau bukan angka
			fmt.Println("Input tidak valid!! Harap masukkan angka yang benar (bilangan bulat)")
			continue
		}
		if val < min || val > max {
			// jika angka diluar range
			fmt.Printf("Input harus di antara %d dan %d.\n", min, max)
			continue
		}
		return val
	}
}

// func yang digunakan Untuk membersihkan layar terminal, membuat tampilan selalu bersih setiap kali ganti menu
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// func digunakan untuk menunggu user, maksudnya menahan layar agar user bisa membaca informasi setiap menu sebelum kembali ke menu utama
func waitForEnter() {
	fmt.Println("\n Tekan [ENTER] untuk kembali ke menu...")
	scanner.Scan()
}
