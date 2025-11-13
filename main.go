package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println("Тут список команд!.....")
}

func main() {

	fmt.Println("Добрый день! Для просмотра команд введите help")

	for {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Введите команду: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()
		fieldScan := strings.Fields(text)

		fmt.Println("Вы ввели команду: ", fieldScan[0])

		cmd := fieldScan[0]

		switch cmd {
		case "help":
			help()
		case "exit":
			return
		case "createuser":

			fmt.Println("Если вы хотите создать нового пользователя введите имя, иначе чтобы выйти exit")
			fmt.Print("Введите имя: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Ошибка ввода!")
				return
			}
			name := scanner.Text()

			if name == "exit" {
				return
			}

			fmt.Println("=======================================")
			fmt.Println("Создание нового пользователя", name, " ...")
			fmt.Println("=======================================")
		default:
			fmt.Println("Вы ничего не ввели! Для просмотра введите 'help'. Повторите ввод...")
		}
	}
}
