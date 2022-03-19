package subject

import "github.com/michelo851a1203/testgo/observer"

type Subject interface {
	NotifyAll()
	AddObserver(obs observer.Observer)
	RemoveObserver(obs observer.Observer)
}
