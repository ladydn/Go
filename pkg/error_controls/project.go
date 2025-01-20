package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	//serializar a json
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil

}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	//serializar a slice
	decoder := json.NewDecoder(file)
	err = decoder.Decode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contact
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Printf("Error loading contacts")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== CONTACT MANAGER ===\n",
			"1. Agregar contacto\n",
			"2. Listar contactos\n",
			"3. Salir\n",
			"4. Elige una opcion: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error leyendo la opcion")
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			contacts = append(contacts, c)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error guardando el contacto")
			}

		case 2:
			fmt.Println("=========================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=========================================================")

		case 3:
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}
