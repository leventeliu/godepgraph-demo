package main

import (
	"fmt"

	"github.com/leventeliu/godepgraph-demo/a"
	"github.com/leventeliu/godepgraph-demo/b"
)

func main() {
	fmt.Println("deps:", a.Name, b.Name)
}
