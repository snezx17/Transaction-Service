package main

import (
	"fmt"
	"net/http"
)

func getbonus(w http.ResponseWriter, r *http.Request) {
	text := "Бонус получен!"
	b := []byte(text)

	_, err := w.Write(b)

	if err != nil {
		fmt.Println("Ошибка отправки")
	} else {
		fmt.Println("Успешная отправка")
	}
}

func main() {
	http.HandleFunc("/bonus", getbonus)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка")
	}
}
