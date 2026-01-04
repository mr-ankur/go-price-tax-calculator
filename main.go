package main

import (
	"fmt"

	"example.com/price-tax-calculator/filemanager"
	"example.com/price-tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
