package main

import (
	"fmt"
	"os"

	"github.com/mlletin/go-torrent/internal/pkg/bencode"
)

func main() {
	file, err := os.Open("sample.torrent")
	if err != nil {
		fmt.Println("failed to open file:", err)
		os.Exit(1)
	}
	result, err := bencode.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
