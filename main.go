package main

import (
	"fmt"

	"github.com/brianseitel/jason/jason"
	"github.com/sanity-io/litter"
)

func main() {
	actual := jason.Lex(`{"foo": "bar", "baz": "beans", "sexy": ["anal", "ass", "tits"]}`)
	fmt.Println("lexed: ", actual)

	litter.Dump(actual)

	out, out2 := jason.Parse(actual)

	litter.Dump(out)
	litter.Dump(out2)
}
