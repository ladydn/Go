package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

const Pi float32 = 3.1415

const (
	x = 100
	y = 0b1010
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	//funciones
	fmt.Println(hello("Luis Mao"))
	fmt.Println(calc(5, 6))

	sum, mult := calc2(5, 6)
	fmt.Println(sum, mult)

	fmt.Println("Hola Mundo")

	firstName, lastName := "Lady", "Diana"
	age := 27

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)

	//tipos de datos basicos
	var integer int8 = 127
	var valoresPositivos uint = 78
	var valoresDecimales float32
	fmt.Println(integer, valoresDecimales, valoresPositivos)

	fmt.Println(math.MinInt64, math.MaxInt64)

	fullName := "Lady Diana \t(alias \"ldcode97\")\n"
	fmt.Println(fullName)

	//valor de codigo ascii
	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0]) //valor decimal de h

	var r rune = '❤'
	fmt.Println(r)

	//valores prederminados
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	//conversiones de tipos
	var interger16 int16 = 50
	var interger32 int32 = 100
	fmt.Println(int32(interger16) + interger32)

	strCon := "100"
	i, _ := strconv.Atoi(strCon)
	fmt.Println(i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)

	//fmt
	nameL := "Lady"
	ageL := 27
	fmt.Printf("Hola me llamo %s y tengo %d anios.\n", nameL, ageL)

	greeting := fmt.Sprintf("Hola me llamo %s y tengo %d anios.\n", nameL, ageL)
	fmt.Println(greeting)

	var nameL2 string
	var ageL2 int
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nameL2)
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d\n", &ageL2)
	fmt.Printf("Hola me llamo %s y tengo %d anios.\n", nameL2, ageL2)

	aa := 10
	bb := 3

	bb--
	bb--
	fmt.Println(aa + bb)

	aaa := 10
	aaa += 20
	fmt.Println(aaa)

	//Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
	var base, altura, area, perimetro float32
	fmt.Print("Ingresa la base del triangulo: ")
	fmt.Scanf("%f\n", &base)
	fmt.Print("Ingresa la altura del triangulo: ")
	fmt.Scanf("%f\n", &altura)
	area = base * altura / 2
	perimetro = base * 3
	fmt.Printf("El area del triangulo es: %f y el perimetro es: %f\n", area, perimetro)

	// control de flujo if, else if, else
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Buenos dias")
	} else if hora < 18 {
		fmt.Println("Buenas tardes")
	} else {
		fmt.Println("Buenas noches")
	}

	//switch os :=runtime.GOOS, os
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Ejecutando en Windows")
	case "linux":
		fmt.Println("Ejecutando en Linux")
	case "darwin":
		fmt.Println("Ejecutando en MacOS")
	default:
		fmt.Println("Go run -> Otros OS")
	}

	//for
	for ii := 0; ii < 5; ii++ {
		fmt.Println(ii)
	}

	for iii := 0; iii < 5; iii++ {
		fmt.Println(iii)
		if iii == 3 {
			break
		}
	}

	for iiii := 0; iiii < 5; iiii++ {
		if iiii == 3 {
			continue
		}
		fmt.Println(iiii)
	}
}

func hello(name string) string {
	return "Hola " + name
}

func calc(a, b int) (int, int) {
	return a + b, a * b
}

func calc2(a, b int) (sum, mult int) {
	sum = a + b
	mult = a * b
	return
}
