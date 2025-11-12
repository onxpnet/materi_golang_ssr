package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

// Struct data untuk dikirim ke templat
type PageData struct {
	Title string // Akan digunakan oleh header.html
	Nama  string // Akan digunakan oleh index.html atau profile.html
}

// Variabel untuk menyimpan templat yang sudah di-parse
var tmpl *template.Template

func init() {
	// KUNCI UTAMA:
	// ParseGlob mengambil SEMUA file yang cocok dengan pola.
	// Ini memuat "index.html", "profile.html", "header.html", dan "footer.html"
	// ke dalam satu koleksi templat.
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	loadEnvFile(".env")
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// Siapkan data untuk halaman ini
	data := PageData{
		Title: "Halaman Utama",
		Nama:  "Staf ESDM",
	}

	// Eksekusi HANYA templat "index.html".
	// Go akan otomatis menemukan "header" dan "footer" karena sudah di-parse.
	err := tmpl.ExecuteTemplate(res, "index.html", data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func profileHandler(res http.ResponseWriter, req *http.Request) {
	// Siapkan data untuk halaman ini
	data := PageData{
		Title: "Profil Pengguna",
		Nama:  "Staf ESDM",
	}

	// Eksekusi templat "profile.html"
	err := tmpl.ExecuteTemplate(res, "profile.html", data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	my_env := GetEnv("MY_ENV", "ini default value")
	fmt.Println(my_env)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/profile", profileHandler)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
