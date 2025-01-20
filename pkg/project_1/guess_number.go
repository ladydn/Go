package main

import (
	"fmt"
	"math/rand"
)

func main() {

	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 5
	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un número: (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scan(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("Ganaste!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")
		}
	}

	fmt.Println("Perdiste! El número a adivinar era: ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var respuesta string
	fmt.Print("Desea jugar nuevamente? (s/n): ")
	fmt.Scan(&respuesta)
	switch respuesta {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("eleccion no válida")
		jugarNuevamente()
	}
}
