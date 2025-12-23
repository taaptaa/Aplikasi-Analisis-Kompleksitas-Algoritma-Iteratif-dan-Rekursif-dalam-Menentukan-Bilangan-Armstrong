package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"runtime"
	"time"
)

// Struktur untuk request JSON dari frontend
type Request struct {
	Number int    `json:"number"`
	Method string `json:"method"`
}

// Struktur untuk response JSON ke frontend
type Response struct {
	Armstrong bool   `json:"armstrong"`
	Time      int64  `json:"time"`
	Memory    uint64 `json:"memory"`
}

func main() {

	// 1. Serve semua file statis (HTML, CSS, JS) dari folder ./static
	// Ini memungkinkan browser memuat index.html, style.css, dan script.js
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// 2. API route untuk logika pengecekan
	http.HandleFunc("/check", checkHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler untuk endpoint /check
func checkHandler(w http.ResponseWriter, r *http.Request) {
	// Pastikan hanya menerima POST
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Pengukuran Kinerja Dimulai
	start := time.Now()
	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	var result bool
	if req.Method == "iterative" {
		result = isArmstrongIter(req.Number)
	} else {
		// Default atau jika method=recursive
		result = isArmstrongRec(req.Number)
	}

	// Pengukuran Kinerja Selesai
	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)

	resp := Response{
		Armstrong: result,
		Time:      time.Since(start).Microseconds(),
		// Menghitung alokasi memori yang terjadi selama eksekusi
		Memory:    memEnd.Alloc - memStart.Alloc,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// --- FUNGSI LOGIKA ARMSTRONG ---

// ITERATIF: Lebih mudah dipahami
func isArmstrongIter(n int) bool {
	if n < 0 {
		return false
	}
	sum := 0
	temp := n
	// Menghitung jumlah digit (d)
	digits := len(fmt.Sprint(n))

	for temp > 0 {
		rem := temp % 10 // Ambil digit terakhir
		// Pangkatkan digit dengan jumlah digit (d)
		sum += int(math.Pow(float64(rem), float64(digits)))
		temp /= 10 // Buang digit terakhir
	}

	return sum == n
}

// REKURSIF: Lebih ringkas namun lebih kompleks dalam stack call
func armstrongHelper(n, digits int) int {
	if n == 0 {
		return 0 // Base case: jika bilangan habis
	}
	// Perhitungan pangkat digit terakhir + hasil rekursi untuk sisa bilangan
	return int(math.Pow(float64(n%10), float64(digits))) + armstrongHelper(n/10, digits)
}

func isArmstrongRec(n int) bool {
	if n < 0 {
		return false
	}
	// Menghitung jumlah digit (d)
	digits := len(fmt.Sprint(n))
	return armstrongHelper(n, digits) == n
}