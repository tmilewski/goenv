package main

import "fmt"
import "os"
import "github.com/tmilewski/goenv"

func init() {
	err := goenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println(os.Getenv("HELLO_WORLD"))
	fmt.Println("Your home directory is:", os.Getenv("HOME_DIR"))
}
