package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1) // increments counter
		// making goroutine to a lambda function so that getCurrencyData and wg.Done() are called synchronously
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done() // decrements counter
		}(currency)
	}
	wg.Wait() // run until the counter does not becomes zero
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
