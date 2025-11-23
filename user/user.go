package user

import "fmt"

var listUsers = make([]userStruct, 0) //БД юзеров

type userStruct struct {
	username string
	balance  int
}

func AddNewUser(n string) { // Создание нового пользователя
	newUser := userStruct{
		username: n,
		balance:  0,
	}

	listUsers = append(listUsers, newUser)

	fmt.Println("Создан новый пользователь:", newUser.username, "баланс: ", newUser.balance)

}

func GetUsers() { // Вывод списка всех пользователей
	fmt.Println("Список всех пользователей:", listUsers)
}

func Addmoney(name string, value int) {

}
