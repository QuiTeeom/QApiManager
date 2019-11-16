package main

import (
	"QApiManager/pkg/api-manager"
	"bufio"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	api_manager.OpenDb()

	dir, _ := os.Getwd()
	mux := http.NewServeMux()

	www := http.Dir(dir + string(filepath.Separator) + "www")
	fileHandler := http.FileServer(www)
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		f, _ := www.Open("index.html")
		defer f.Close()

		b := make([]byte, 1024)

		reader := bufio.NewReader(f)
		for {
			i, _ := reader.Read(b)
			if i <= 0 {
				break
			}
			writer.Write(b[0:i])
		}
	})

	mux.Handle("/js/", fileHandler)

	mux.Handle("/modules/", fileHandler)

	api_manager.CreateHandler(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
