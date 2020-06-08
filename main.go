package main

import (
	"fmt"

	"github.com/brianseitel/jason/jason"
	"github.com/sanity-io/litter"
)

func main() {
	actual := jason.Lex(`{"foo": "bar", "baz": "beans", "bears": ["grizzly", "brown", "polar"]}`)
	fmt.Println("lexed: ", actual)

	out, _ := jason.Parse(actual)

	fmt.Println("parsed:")
	litter.Dump(out)
}
