package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
		"quarter_of_a_dozen" : 3,
		"half_of_a_dozen":	6,
		"dozen":	12,
		"small_gross":	120,
		"gross":	144,
		"great_gross":	1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if exists {
		bill[item] += units[unit]
		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	_, itemExists := bill[item]

	fmt.Printf("bill= %v | units= %v | item= %s | unit= %s \n", bill, units, item, unit)

	if unitExists && itemExists {
		newQty := bill[item] - units[unit]
		if newQty == 0 {
			delete(bill, item)
			return true
		} else if newQty < 0 {
			return false
		} 
		bill[item] = newQty
		return true		
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	
	return qty, ok
}
