package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1{
		message := ""
		for i := 1; i < len(os.Args); i++{
			message += os.Args[i]
			if i != len(os.Args)-1{
				message += " "
			}
		}
		fmt.Println("Hello "+ message +". Welcome to the jungle")
	}else{
		fmt.Println("No arguments received")
	}
	
}
