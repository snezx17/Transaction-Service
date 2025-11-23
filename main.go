package main

import (
	"Transaction-Service/user"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Println("=======================================")
	fmt.Println("Список команд:")
	fmt.Println("users | Вывести список всех пользователей")
	fmt.Println("createuser | Создать нового пользователя")

	fmt.Println("=======================================")

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

		case "help": // Список команд
			help()

		case "exit": // Выход
			return

		case "users":
			user.GetUsers()

		case "createuser": // Создание пользователя

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

			if len(name) >= 2 {
				fmt.Println("=======================================")
				fmt.Println("Создание нового пользователя", name, " ...")
				fmt.Println("=======================================")

				user.AddNewUser(name)

				fmt.Println("Новый пользователь создан!")

			} else {
				fmt.Println("Вы ввели некорректное имя!")
			}
		case "addmoney": // Добавление balance к user
			fmt.Print("Укажите имя пользователя для начисления: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Ошибка ввода!")
				return
			}
			name := scanner.Text()
			fmt.Println("Корректно ли имя:", name)

			if ok := scanner.Scan(); !ok {
				fmt.Println("Ошибка ввода!")
				return
			}

			input := scanner.Text()

			if input == "y" {
				if ok := scanner.Scan(); !ok {
					fmt.Println("Ошибка ввода!")
					return
				}
				input := scanner.Text()
				value, err := strconv.Atoi(input)

				if err != nil {
					fmt.Println("Ошибка ввода!")
					return
				}

				user.Addmoney(name, value)

			}
			if input == "n" {

			} else {

			}
			//Вывод пользователя + подтверждение
			//ввод валюты
		case "removemoney": // Удаление balance к user

		case "transfermoney": // Перевод balance от user к user

		default:
			fmt.Println("Вы ничего не ввели! Для просмотра введите 'help'. Повторите ввод...")
		}
	}
}
