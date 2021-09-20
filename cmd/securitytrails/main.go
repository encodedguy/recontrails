package main

import (
	"fmt"
	"os"
	"log"

	"securitytrails"
)

func main() {
	if len(os.Args) != 1{
		log.Fatalln("Usage: ./main")
	}

	apiKey := os.Getenv("SECURITYTRAILS_API_KEY")

	s := securitytrails.New(apiKey)
	info, err := s.APIInfo()
	usage, err := s.APIUsage()

	if err != nil{
		log.Panicln(err)
	}

	fmt.Printf("APIKey Valid: %t\n", info.Success)
	fmt.Printf("API Allowed Credits: %d\nAPI Used Credits: %d\n", usage.AllowedMonthlyUsage, usage.CurrentMonthlyUsage)
}
