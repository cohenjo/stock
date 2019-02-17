package scanner

import (
	"fmt"
)

func Scan() {

	fmt.Printf("#############################################################################################################\n")
	cs := Fish()
	fmt.Printf("#############################################################################################################\n")
	fmt.Printf("#############################################################################################################\n")
	// RegisterFilters()
	RegisterHighFilters()

	cs = FR.RunFilters(cs)

	for _, s := range cs {
		fmt.Printf("found: %s, Ticker: %s, Yield: %f chow: %f\n", s.Name, s.Ticker, s.DividendYield, s.Chowder)
	}
	fmt.Printf("#############################################################################################################\n")

	// symbols := make([]string, len(cs))
	// for i, s := range cs {
	// 	symbols[i] = s.Ticker
	// }

	// iter := equity.List(symbols)

	// // Iterate over results. Will exit upon any error.
	// for iter.Next() {
	// 	q := iter.Equity()
	// 	fmt.Println(q)
	// }
	fmt.Printf("#############################################################################################################\n")
}

// look into: https://github.com/aktau/gofinance/blob/master/app/main.go
