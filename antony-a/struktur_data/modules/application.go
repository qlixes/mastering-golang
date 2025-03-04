package modules

import (
	"fmt"
)

func Exchange(billing int, payment int) {
	currency := []int{50000, 20000, 10000, 5000, 2000, 1000, 500, 100}
	balance := payment - billing

	fmt.Printf("Your exchange was %d with detail : \n", balance)

	result := make(map[int]int)

	for _, money := range currency {
		if balance > 100 {
			post := balance / money
			balance = balance - (post * money)
			result[money] = post
		}
	}

	fmt.Println(result)
}
