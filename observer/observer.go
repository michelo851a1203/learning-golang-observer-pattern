package observer

type Observer interface {
	GetId() string
	Notify(string)
}
