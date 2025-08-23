package main

import (
	"fmt"

	"github.com/sethll/goCBC/pkg/progmeta"
)

func main() {
	fmt.Print(progmeta.ProgVersion.Tag())
}
