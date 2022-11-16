package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(00),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(30),
		Fruits:              FruitsPercent(60),
		Fiber:               FiberGram(100),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
