package main

import "fmt"

// son la variables que almacenan la direccion de la memoria de otra variable
// son usados para acceder y referenciar a una variable original a travez de su direccion de la memoria

type Persona2 struct {
	nombre string
	edad   int
	correo string
}

// receptor
func (p *Persona2) sayHello() {
	fmt.Println("Hola", p.nombre)
}

func edit(x *int) {
	*x = 20
}

func main() {
	//var x int = 10

	//var p *int = &x
	//fmt.Println(&x)
	//fmt.Println(p)

	//fmt.Println(x)
	//edit(&x)
	//fmt.Println(x)

	p := Persona2{nombre: "Juan", edad: 30, correo: "juan@gmail.com"}
	p.sayHello()
}
