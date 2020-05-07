package main

import (
	"errors"
)

func main() {

}

func concat(a string, b string) (string, error) {
	if a == "" || b == "" {
		return "", errors.New("fail")
	}
	return (a + b), nil
}
