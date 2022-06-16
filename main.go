package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kilogram float64
type Libras float64
type Foot float64
type Meters float64

func main() {
	var option int
	fmt.Println("Informe qual conversão você gostaria de fazer:")
	fmt.Println("1 - Celsius para Farenheit")
	fmt.Println("2 - Fahrenheit para Celsius")
	fmt.Println("3 - Quilograma em Libras")
	fmt.Println("4 - Libras para Quilograma")
	fmt.Println("5 - Pés para Metros")
	fmt.Println("6 - Metros para Pés")
	fmt.Scanln(&option)

	switch option {
	case 1:
		var celsius Celsius
		fmt.Println("Informe o valor em Celsius:")
		fmt.Scanln(&celsius)
		result := celsiusToFahrenheit(celsius)
		fmt.Printf("O resultado é %.2f C", result)
	case 2:
		var fahrenheit Fahrenheit
		fmt.Println("Informe o valor em Fahrenheit:")
		fmt.Scanln(&fahrenheit)
		result := fahrenheitToCelsius(fahrenheit)
		fmt.Printf("O resultado é %.2f F", result)
	case 3:
		var weight Kilogram
		fmt.Println("Informe o valor em Quilogramas:")
		fmt.Scanln(&weight)
		result := kilogramToLibras(weight)
		fmt.Printf("O resultado é %.2f lb", result)
	case 4:
		var weight Libras
		fmt.Println("Informe o valor em Libras")
		fmt.Scanln(&weight)
		result := librasToKilogram(weight)
		fmt.Printf("O resultado é %.2f kg", result)
	case 5:
		var length Foot
		fmt.Println("Informe o valor em Pés:")
		fmt.Scanln(&length)
		result := footToMeters(length)
		fmt.Printf("O resultado é %.2f m", result)
	case 6:
		var length Meters
		fmt.Println("Informe o valor em Metros:")
		fmt.Scanln(&length)
		result := metersToFoot(length)
		fmt.Printf("O resultado é %.2f ft", result)
	default:
		fmt.Println("Opção inválida")
	}
}

func celsiusToFahrenheit(temp Celsius) Fahrenheit {
	return Fahrenheit((temp * 9 / 5) + 32)
}

func fahrenheitToCelsius(temp Fahrenheit) Celsius {
	return Celsius((temp - 32) * 5 / 9)
}

func kilogramToLibras(weight Kilogram) Libras {
	val1 := weight * 2
	val2 := val1 / 10
	return Libras(val1 + val2)
}

func librasToKilogram(weight Libras) Kilogram {
	return Kilogram(weight / 2.2046)
}

func footToMeters(length Foot) Meters {
	return Meters(length * 0.3048)
}

func metersToFoot(length Meters) Foot {
	return Foot(length * 3.2808)
}
