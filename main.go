package main

import (
	"fmt"

	// "example.com/price-tax-calculator/cmdmanager"
	"example.com/price-tax-calculator/filemanager"
	"example.com/price-tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
