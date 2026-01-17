package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var Balance = 1000

func getbonushandler(w http.ResponseWriter, r *http.Request) {
	text := "Бонус получен!"
	b := []byte(text)

	_, err := w.Write(b)

	if err != nil {
		fmt.Println("Ошибка отправки")
	} else {
		fmt.Println("Успешная отправка")
	}
}

func payhanler(w http.ResponseWriter, r *http.Request) {
	ResponsBodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to read http body")
		return
	}

	ResponsBodyRequestString := string(ResponsBodyRequest)

	paymentAmmount, err := strconv.Atoi(ResponsBodyRequestString)
	if err != nil {
		fmt.Println("failed to convert http body to integer")
	}

	if paymentAmmount > Balance {
		msg := "Нехватка баланса" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write http response")
			return
		}
	} else {
		Balance -= paymentAmmount
		msg := "Успешное списание"
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write http response")
			return
		}

	}
}

func main() {
	http.HandleFunc("/bonus", getbonushandler)
	http.HandleFunc("/pay", payhanler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Произошла ошибка")
		return
	}
}
