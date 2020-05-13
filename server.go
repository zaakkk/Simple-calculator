package main

import "net/http"

func main() { // メイン処理
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
