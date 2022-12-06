package main

import "fmt"
import "os"

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please give us something")
		os.Exit(1)
	}

	file := arguments[1]

	info, err := os.Stat(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mode := info.Mode()
	fmt.Print(file, ": ", mode, "\n")
}
