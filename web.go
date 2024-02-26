package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var htmlStr string

func main() {
	fmt.Println("Starting the server!")

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	htmlStr = string(data)

	// ルートとハンドラ関数を定義
	http.HandleFunc("/", showHtml)
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/categories", categoryHandler)
	http.HandleFunc("/api/calculator", calculateHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	name := query.Get("name")

	// レスポンス用マップを作成
	response := map[string]string{
		"message": "Hello " + name,
	}

	// Content-typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをjsonにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	var animals []string = []string{"dog", "cat", "bird", "fox"}

	response := map[string][]string{
		"category": animals,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	operator := query.Get("o")

	x, err := strconv.ParseFloat(query.Get("x"), 64)
	if err != nil {
		http.Error(w, "x is not number. x must be number,", http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseFloat(query.Get("y"), 64)
	if err != nil {
		http.Error(w, "y is not number. y must be number,", http.StatusBadRequest)
		return
	}

	var result float64
	switch operator {
	case " ":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			http.Error(w, "error: division by 0.", http.StatusBadRequest)
			return
		}
		result = x / y
	default:
		http.Error(w, "Invalid operator", http.StatusBadRequest)
		return
	}

	response := map[string]float64{
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func showHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlStr)
}
