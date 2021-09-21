package main

import (
	"fmt"
	"os"
	"log"
	"securitytrails"
)

func main() {
	if len(os.Args) != 2{
		log.Fatalln("Usage: ./main domain")
	}

	apiKey := os.Getenv("SECURITYTRAILS_API_KEY")

	s := securitytrails.New(apiKey)

	info, err := s.APIInfo()
	if err != nil{
		log.Panicln(err)
	}

	usage, err := s.APIUsage()
	if err != nil{
		log.Panicln(err)
	}

	dnsa, err := s.DnsA(os.Args[1])
	if err != nil{
		log.Panicln(err)
	}

	fmt.Printf("APIKey Valid: %t\n", info.Success)
	fmt.Printf("API Allowed Credits: %d\nAPI Used Credits: %d\n", usage.AllowedMonthlyUsage, usage.CurrentMonthlyUsage)
	for _, value := range dnsa.Organizations{
		fmt.Printf("%s\n", value)
	}
	fmt.Printf("%s", dnsa.Organizations[1])
}
