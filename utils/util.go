package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = "abcdefghijklmnopqrstuvxyz"

func generate_string(lenth int) (ans string) {

	for len(ans) < lenth {
		ans += string(letters[rand.Intn(len(letters))])
	}
	return
}

func generate_int() string {
	t := fmt.Sprint(time.Now().Nanosecond())
	tm := t[:7]
	return tm
}

func RandomUserName(lenth int) string {
	return generate_string(lenth)
}

func RandomUserContact() string {
	return generate_int()
}

func RandomUserEmail() string {
	user_name := generate_string(10)
	return fmt.Sprintf("%s@gmail.com", user_name)

}
