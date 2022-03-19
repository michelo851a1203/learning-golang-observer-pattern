package main

import (
	"github.com/michelo851a1203/testgo/customer"
	"github.com/michelo851a1203/testgo/product"
)

// what is observer pattern ? :
// ===== observer pattern =====
// if subject state update and then notify all observer
// application situation : if product is on sale then notify all customer

func main() {
	// build subject first
	productSubject := product.BuildItemInstanceWithName("testing")
	// observer here
	customerOne := customer.Customer{ID: "observer testing 1"}
	customerTwo := customer.Customer{ID: "observer testing 2"}

	productSubject.AddObserver(&customerOne)
	productSubject.AddObserver(&customerTwo)

	productSubject.UpdateAvailability()
}
