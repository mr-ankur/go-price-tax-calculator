package main

import (
	"fmt"

	// "example.com/price-tax-calculator/cmdmanager"
	"example.com/price-tax-calculator/filemanager"
	"example.com/price-tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Cound not process job")
			fmt.Println(err)
		}
	}
}
