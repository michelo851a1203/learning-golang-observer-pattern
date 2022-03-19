package product

import "github.com/michelo851a1203/testgo/observer"

type Product struct {
	ObserverList []observer.Observer
	Name         string
	IsInStock    bool
}

func BuildItemInstanceWithName(name string) *Product {
	return &Product{
		Name: name,
	}
}

func (product *Product) NotifyAll() {
	for _, observer := range product.ObserverList {
		observer.Notify(product.Name)
	}
}

func (product *Product) AddObserver(obs observer.Observer) {
	product.ObserverList = append(product.ObserverList, obs)
}

func (product *Product) RemoveObserver(obs observer.Observer) {
	product.ObserverList = removeFromSlice(product.ObserverList, obs)
}

func removeFromSlice(observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
	observerLength := len(observerList)
	for index, observer := range observerList {
		if observer.GetId() == observerToRemove.GetId() {
			observerList[observerLength-1], observerList[index] = observerList[index], observerList[observerLength-1]
			return observerList[:observerLength-1]
		}
	}
	return observerList
}

func (product *Product) UpdateAvailability() {
	product.IsInStock = true
	product.NotifyAll()
}
