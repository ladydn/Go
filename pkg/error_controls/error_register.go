package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main(): ")
	//log.Print("Oye, soy un log")
	//log.Fatal("This is a fatal message")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Oye, soy un log")
}
