package util

import "errors"

const HelloworldMessage = "hello eorld from constant"
//public function alway have to start with uppercase
func GetMessage() (string, int) {

	message := "Hello world"

	number := 1
	return message, number

}

func CreateMessage(x int) (string, int, error) {

	message := "Hello Krish"
	if x < 0 {
		return "", 0, errors.New("x cannot be less then 0")
	}
	number := x
	return message, number, nil

}