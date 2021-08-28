package controllers

import (
	"net/http"

	"github.com/soutaschool/todo-app/config"
)

func StartMainServer() error {
	// 通常もnilを入れる
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
