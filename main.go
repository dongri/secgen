package main

import (
	"flag"
	"fmt"
	"math/big"

	"crypto/rand"
	"encoding/base64"
)

var (
	optLength   = flag.Int("l", 10, "\"i\" length")
	optQuantity = flag.Int("q", 1, "\"q\" quantity")
	optEncoding = flag.String("e", "", "\"e\" encoding (ex: base64)")
)

func main() {
	flag.Parse()

	length := *optLength
	qty := *optQuantity
	encoding := *optEncoding

	for i := 0; i < qty; i++ {
		secret, err := generateRandomString(length)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		if encoding == "base64" {
			secret = base64.StdEncoding.EncodeToString([]byte(secret))
		}
		fmt.Println(secret)
	}
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
