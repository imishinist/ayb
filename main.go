package main

import (
	"fmt"

	"github.com/imishinist/ayb/pkg/witticism"
)

func main() {
	witticism := witticism.Get()
	fmt.Println(witticism.Random().Text)
}
