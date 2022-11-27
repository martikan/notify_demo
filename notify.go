package main

type Observer interface {
	update(string)
	getID() string
}

type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}
