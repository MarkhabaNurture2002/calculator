package main

import "fmt"

const BaseRate = 5.50
const TaxRate = 0.12
const DistanceRate = 2.0
const FragileFee = 0.2

func main() {
	var name string
	var weight float64
	var distance float64
	var fragileCount int

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Вес груза (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Дистанция (км): ")
	fmt.Scan(&distance)

	fmt.Print("Количество хрупких упаковок: ")
	fmt.Scan(&fragileCount)

	var calculation = (weight * BaseRate) * (1 + FragileFee*float64(fragileCount) + (distance * DistanceRate))

	tax := calculation * TaxRate
	totalCost := calculation + tax

	fmt.Println("--- Отчет о доставке ---")
	fmt.Printf("Отправитель: %s\n", name)
	fmt.Printf("Итоговая стоимость: %.2f\n", totalCost)

}
