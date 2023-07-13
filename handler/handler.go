package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//membuat kondisi dimana jika routing tidak ada atau belum dibuat maka tidak menampilkan halaman Home/Root
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		w.Write([]byte("BAPAK KAU BOTAK"))
		return

	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html")) // membuat agar web dinamis saat menambahkan layout

	if err != nil {
		log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
		http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
		return
	}
	// AKSES DATA MAP
	// data := map[string]interface{}{
	// 	"title":   "Belajar With Amau",
	// 	"content": "Golang Web season 1",
	// }

	//AKSES DATA STRUCT DI ENTITY.GO
	// data := entity.Product{ID: 1, Name: "Amau", Stock: 4}

	// Akses menggunakan slice
	data := []entity.Product{
		{ID: 1, Name: "Buah", Stock: 3},
		{ID: 2, Name: "Rumput", Stock: 9},
		{ID: 3, Name: "Motor", Stock: 12},
	}
	err = tmpl.Execute(w, data) //untuk menampilkan halaman html nya

	if err != nil {
		log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
		http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
		return
	}
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Aku Mario"))
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Aku Hello jir"))
}

// mengambil Query string contoh : root/product?id=2

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// karena id masih berupa string maka cara mengecek nya dengan strconv.Atoi yang mengkonversikan nya
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
		http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
		return
	}

	data := map[string]interface{}{
		"idProduct": idNumb,
	}
	err = tmpl.Execute(w, data) //untuk menampilkan halaman html nya

	if err != nil {
		log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
		http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
		return
	}

}

func PostGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("Post"))
	case "GET":
		w.Write([]byte("GET"))
	default:
		http.Error(w, "tenang", http.StatusBadRequest)
	}
}

func RouteIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
			http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
			http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
			return
		}

		return
	}
	http.Error(w, "keep calm", http.StatusBadRequest)
}

func Result(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
			http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
			return
		}

		name := r.Form.Get("name")
		msg := r.Form.Get("message")

		data := map[string]interface{}{
			"name": name,
			"msg":  msg,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)                                                      // pesan error untuk developer yang akan di tampilan di terminal
			http.Error(w, "Error lagi bro bruhh", http.StatusInternalServerError) //Pesan error untuk user
			return
		}

		err = tmpl.Execute(w,data)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
