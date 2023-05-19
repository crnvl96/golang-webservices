package main

import (
	"fmt"
	"io"
	"os"

	"github.com/crnvl96/golang-webserver/helpers"
)

func main() {
	ctx, cancel := helpers.GenerateContextWithTimeout(300)
	defer cancel()

	res, err := helpers.ExecuteNewRequestWithContext(ctx, "http://localhost:8080/cotacao")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
