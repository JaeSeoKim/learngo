package main

import (
	"fmt"

	"github.com/JaeSeoKim/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "first", "second": "second", "third": "third"}
	dictionary.Delete("h")
	fmt.Println(dictionary.Search("h"))
}
