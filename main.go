package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"crypto/rand"
)

const (
	DefaultLength = 10
	DefaultQty    = 1
)

func main() {
	var length int
	var qty int
	var err error
	if len(os.Args) < 2 {
		fmt.Println("Please set length and quantity")
		fmt.Println("usage: secgen <n: length> [option] <n: quantity>")
		return
	} else {
		length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		qty = DefaultQty
		if len(os.Args) > 2 {
			qty, err = strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Print(err.Error())
				return
			}
		}
	}

	for i := 0; i < qty; i++ {
		secret, err := generateRandomString(length)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Println(secret)
	}
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+|~-=[]{};:"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}
