package main

import (
	"fmt"

	"github.com/fyntrix/fyntrix/image/vips"
)

func main() {
	vip := vips.New()
	fmt.Println(vip.Version())
}
