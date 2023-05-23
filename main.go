package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/workspace/{workspaceid}/user/{userid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["workspaceid"]
		page := vars["userid"]

		fmt.Fprintf(w, "あなたの所属ワークスペース名は %s で、あなたは %s さんです。", title, page)
	})

	// 静的ファイルの読み込み
	fs := http.FileServer(http.Dir("static/"))
	// /static/にアクセスがあれば、/staticの末尾にfsに格納されているファイル名を付与
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}
