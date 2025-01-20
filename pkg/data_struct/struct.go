package main

import "fmt"

// el modelo de una clase
type Persona struct {
	Nombre string
	Edad   int
	Correo string
}

func main() {
	var p Persona
	p.Nombre = "Juan"
	p.Edad = 30
	p.Correo = "juan@gmail.com"

	//instanciar un objeto
	p2 := Persona{Nombre: "Pedro", Edad: 25, Correo: "juan@gmail.com"}
	p2.Edad = 26

	fmt.Printf("Nombre: %s, Edad: %d, Correo: %s\n", p.Nombre, p.Edad, p.Correo)
	fmt.Println(p2)
}
