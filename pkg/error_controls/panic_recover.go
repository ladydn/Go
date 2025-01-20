package main

import "fmt"

func dividePanic(dividendo, divisor int) {
	defer func() {
		//el recover recupera el panico, este panico puede ser personalizado
		//o puede mostrar el panico original
		if r := recover(); r != nil {
			fmt.Printf("Recuperado en dividePanic: %v\n", r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
}

// el panico interrumpe la ejecucion ante un error
func main() {
	dividePanic(100, 10)
	dividePanic(80, 10)
	dividePanic(100, 0)
	dividePanic(64, 8)
}
