package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	// HTTPサーバーが稼働中か確認するための/healthエンドポイントを宣言してある
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
