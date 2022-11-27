package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	available    bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)

	i.available = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	ch := make(chan Observer)
	go func() {
		for _, o := range i.observerList {
			ch <- o
		}

		close(ch)
	}()

	for v := range ch {
		v.update(i.name)
	}

	fmt.Println("all customers have been notified")
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)

	for i, o := range observerList {
		if observerToRemove.getID() == o.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}

	return observerList
}
