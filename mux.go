package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ryo0210/go_todo_app/handler"
	"github.com/ryo0210/go_todo_app/store"
)

func NewMux() http.Handler {
	// multiplexer（マルチプレクサ）
	mux := chi.NewRouter()
	// HTTPサーバーが稼働中か確認するための/healthエンドポイントを宣言してある
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux
}