package main

import (
	"fmt"
	"strings"
)

var productPrice = map[string]float64{
	"T-Shirt": 20.543,
	"MUG":     12.242,
	"HAT":     18,
	"Book":    25.99,
}

func calculateItemPrice(itemName string) (float64, bool) {
	basePrice, found := productPrice[itemName]
	if !found {
		if strings.HasSuffix(itemName, "_SALE") {
			originalItemName := strings.TrimSuffix(itemName, "_SALE")
			basePrice, found = productPrice[originalItemName]
			
			if found {
				salePrice := basePrice * .90
				fmt.Printf("Item %s (Orginal Price : $%.2f, Sale Price: $%.2f)\n",
					itemName, basePrice, salePrice)
				return salePrice, true
			}

		} else {
			fmt.Printf("Item: %s (Product not found)\n", itemName)
			return 0, false	
		}
	}
	return basePrice, found
}

func main() {

	orderItems := [] string {
		"T-Shirt", "MUG_SALE", "HAT", "Book_SALE", "BAT",
	}

	var subTotal float64

	for _ , items := range orderItems {
		price, isAvilable := calculateItemPrice(items)
		if isAvilable {
			subTotal += price
			fmt.Printf("Item is avilable at price : $%.2f\n", price)
		}else{
			fmt.Println("Item Out of stock");
		}
	}

	fmt.Printf("Your total amount is : %.2f\n", subTotal);

}