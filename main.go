package main

import (
	"fmt"
)

func main() {
	fmt.Println("Приветствую Вас в регистрационной форме! v.1.0")
	fmt.Println("Чтобы начать регистрацию, введите свой логин:")

	var username string
	fmt.Scan(&username)

	fmt.Println("Теперь придумайте пароль:")
	var password string
	fmt.Scan(&password)

	fmt.Println("Сколько Вам лет?")
	var age int
	fmt.Scan(&age)

	fmt.Print("Поздравляю, ", username, ", теперь вы зарегистрированы!\n")
	fmt.Print("Ваш пароль: ", password, "\n")
	fmt.Print("Ваш возраст: ", age, "\n")

}
