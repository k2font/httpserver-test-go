package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ユーザー情報を取得する処理
// 今回はMock
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workspace := vars["workspaceid"]
	user := vars["userid"]

	fmt.Fprintf(w, "あなたの所属ワークスペース名は %s で、あなたは %s さんです。", workspace, user)
}

// ユーザーを作成する処理
// 今回はMock
func createUser(w http.ResponseWriter, r *http.Request) {

}

// ユーザー情報を更新する処理
// 今回はMock
func updateUser(w http.ResponseWriter, r *http.Request) {

}

// ユーザー情報を削除する処理
// 今回はMock
func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/workspace/{workspaceid}/user/{userid}", getUser).Methods("GET")
	r.HandleFunc("/workspace/{workspaceid}/user/{userid}", createUser).Methods("POST")
	r.HandleFunc("/workspace/{workspaceid}/user/{userid}", updateUser).Methods("PATCH")
	r.HandleFunc("/workspace/{workspaceid}/user/{userid}", deleteUser).Methods("DELETE")

	// 静的ファイルの読み込み
	fs := http.FileServer(http.Dir("static/"))
	// /static/にアクセスがあれば、/staticの末尾にfsに格納されているファイル名を付与
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}
