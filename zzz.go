package main

import (
	"fmt"
	"os"
	"time"
)

const (
	BLUE = "\033[34m"
	END  = "\033[0m"
)

func main() {
	fmt.Fprint(os.Stdout, BLUE)
	for i := 0; i < 5; i++ {
		for _, zzz := range zzzs {
			fmt.Fprintf(os.Stdout, "\r%s", zzz)
			time.Sleep(500 * time.Millisecond)
			fmt.Fprint(os.Stdout, "\033[13A")
		}
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprint(os.Stdout, "\033[1A")
	}
	fmt.Fprint(os.Stdout, END)
}
