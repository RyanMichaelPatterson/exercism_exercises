package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string] int {}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    unitValue, unitExists := units[unit]
    if unitExists{
		billValue, billExists := bill[item]
		if billExists{
            bill[item] = unitValue + billValue
        }else{
			bill[item] = unitValue
        }
        return true
    }
    return false
}        

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	billValue, billExists := bill[item]    

    if unitExists && billExists {
        if billValue - unitValue < 0 {
        	return false
    	}
        if billValue == unitValue {
            delete(bill, item)
        } else {
            bill[item] = billValue - unitValue
        }
        return true
    }
    return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billValue, billExists := bill[item]
    if billExists {
        return billValue, true
    }
    return 0, false
}
