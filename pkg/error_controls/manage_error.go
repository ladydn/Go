package main

import (
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//return 0, errors.New("No se puede dividir por cero")
		return 0, fmt.Errorf("No se puede dividir por cero")
	}
	result := dividendo / divisor
	return result, nil
}

func main() {
	//str := "123"
	//num, err := strconv.Atoi(str)
	//if err != nil {
	//	fmt.Println("Error al convertir el número", err)
	//	return
	//}
	//fmt.Println("Número convertido:", num)

	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error al dividir:", err)
		return
	}
	fmt.Println("Resultado:", result)
}
