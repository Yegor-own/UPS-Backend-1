package models

import "fmt"

type Alphabet []rune

func GetAlphabet() Alphabet {
	return []rune("abcdefghijklmnopqrstuvwxyz")
}

func GetRuneId(char rune) int {
	for i, v := range GetAlphabet() {
		if v == char {
			return i
		}
	}
	return -1
}

func EncodeAlphabet(msg string, rot int) string {
	var res string
	for _, char := range msg {
		charId := GetRuneId(char)
		if charId < 0 {
			fmt.Println("GetRuneId >> returned an error")
			return ""
		}
		res += string(GetAlphabet()[(charId+rot)%26])
	}
	return res
}

func DecodeAlphabet(msg string, rot int) string {
	var res string
	for _, char := range msg {
		charId := GetRuneId(char)
		if charId < 0 {
			fmt.Println("GetRuneId >> returned an error")
			return ""
		}
		res += string(GetAlphabet()[(charId-rot)%26])
	}
	return res
}
