package main

import (
	"desc/pkg/describe"
	"fmt"
)

func main() {

	//бмв
	bmw := describe.NewAuto(describe.CM, "BMW", "Х5", 492.2, 200.4, 174.5, 250, 625)
	fmt.Println("\nБрэнд:", bmw.Brand(), "Модель:", bmw.Model(), "Мощность двигателя:", bmw.EnginePower(), "Максимальная скорость:", bmw.MaxSpeed())
	bmwDimensions := bmw.Dimensions()
	fmt.Println("Высота:", bmwDimensions.Height().Get(describe.CM), "Ширина:", bmwDimensions.Width().Get(describe.CM), "Длина:", bmwDimensions.Length().Get(describe.CM))

	//мерседес
	mercedes := describe.NewAuto(describe.CM, "Mercedes", "Maybach", 520.5, 203, 183.8, 350, 530)
	fmt.Println("\nБрэнд:", mercedes.Brand(), "Модель:", mercedes.Model(), "Мощность двигателя:", mercedes.EnginePower(), "Максимальная скорость:", mercedes.MaxSpeed())
	mercedesDimensions := mercedes.Dimensions()
	fmt.Println("Высота:", mercedesDimensions.Height().Get(describe.CM), "Ширина:", mercedesDimensions.Width().Get(describe.CM), "Длина:", mercedesDimensions.Length().Get(describe.CM))

	//додж
	dodge := describe.NewAuto(describe.CM, "Dodge", "Ram", 581.6, 201.6, 192.3, 190, 702)
	fmt.Println("\nБрэнд:", dodge.Brand(), "Модель:", dodge.Model(), "Мощность двигателя:", dodge.EnginePower(), "Максимальная скорость:", dodge.MaxSpeed())
	dodgeDimensions := dodge.Dimensions()
	fmt.Println("Высота:", dodgeDimensions.Height().Get(describe.Inch), "Ширина:", dodgeDimensions.Width().Get(describe.Inch), "Длина:", dodgeDimensions.Length().Get(describe.Inch))
}
