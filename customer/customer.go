package customer

import "fmt"

// observer

type Customer struct {
	ID string
}

func (customer *Customer) GetId() string {
	return customer.ID
}

func (customer *Customer) Notify(itemName string) {
	fmt.Printf("send product info to customer : %s - notify : %s\n", customer.ID, itemName)
}
