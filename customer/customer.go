package customer

import "fmt"

// observer

type Customer struct {
	id string
}

func (customer *Customer) GetId() string {
	return customer.id
}

func (customer *Customer) Notify(itemName string) {
	fmt.Printf("send product info to customer : %s - %s\n", customer.id, itemName)
}
