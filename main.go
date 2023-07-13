package main

import (
	h "golangweb/handler"
	"log"
	"net/http"
)

func main() {
	// Membuat routing

	mux := http.NewServeMux()

	// bisa juga membuat routing dengan closure yakni membuat fungsi di dalam var
	about := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hallo aku about page"))
	}

	// jika routing yang dipanggil salah satu  maka code yang dibawah ini akan dijalankan
	mux.HandleFunc("/", h.HomeHandler) //membuat halaman homepage atau sering disebut root
	mux.HandleFunc("/hello", h.HelloHandler)
	mux.HandleFunc("/mario", h.MarioHandler)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/product", h.ProductHandler)
	mux.HandleFunc("/hayu", h.PostGet)
	mux.HandleFunc("/process", h.RouteIndexGet)
	mux.HandleFunc("/result", h.Result)

	// bisa juga menggunakan anonymous function
	mux.HandleFunc("/profil", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello profil"))

	})

	// Cara memanggil CSS
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// menampilkan pesan ketika berhasil berjalan
	log.Println("Starting web on port 8080")

	// untuk menjalankan server nya bisa dengan

	err := http.ListenAndServe(":8080", mux)

	// jika ada error maka bisa menggunakan fatal agar langsung berhenti ketika ada error
	log.Fatal(err)
}
