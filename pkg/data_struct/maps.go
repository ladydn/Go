package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
	}

	fmt.Println(colors["red"])
	colors["black"] = "#000000"
	fmt.Println(colors)

	valor, ok := colors["black"]
	if ok {
		fmt.Println("si existe")
	} else {
		fmt.Println("no existe")
	}

	fmt.Println(valor, ok)

	if valor2, ok2 := colors["white"]; ok2 {
		fmt.Println(valor2)
	} else {
		fmt.Println("no existe")
	}

	//eliminar un elemento
	delete(colors, "red")
	fmt.Println(colors)

	for key, valur := range colors {
		fmt.Printf("key: %s, value: %s\n", key, valur)
	}
}
