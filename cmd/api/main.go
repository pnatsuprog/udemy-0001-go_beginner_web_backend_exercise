package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler は、"/hello" に対するハンドラです。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// healthCheckHandler は、"/health" に対するハンドラです。
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// ハンドラの設定
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthCheckHandler)

	// サーバーをポート8080で起動
	port := "8080"
	fmt.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
