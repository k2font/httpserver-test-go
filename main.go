package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This page is a sample page")
	})

	// 静的ファイルの読み込み
	fs := http.FileServer(http.Dir("static/"))
	// /static/にアクセスがあれば、/staticの末尾にfsに格納されているファイル名を付与
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
