package helper

import "fmt"

func ShowOptions(options []string) {
	for _, option := range options {
		fmt.Println(option)
	}
}
