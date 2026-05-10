package utils

import "fmt"

func Logger(tag string, message string) {
	fmt.Println("[BACKEND::", tag, "] ", message)
}
