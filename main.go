package main 

import (
	"fmt"
	"os"
)

func main() {
	if err := run(), err != nil {
		fmt.FprintIn(os.stderr, err)
	}
}


func run() error {
	retun nil
}