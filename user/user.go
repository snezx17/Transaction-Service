package user

import "fmt"

type User struct {
	name    string
	balance float32
}

func (u User) Createuser(n string) {
	if len(n) < 2 {
		fmt.Println("Вы ввели некорректное имя!")
		return
	}
	u.name = n
	u.balance = 0.0

	fmt.Println("Новый пользователь:", u.name, "баланс: ", u.balance)
}
