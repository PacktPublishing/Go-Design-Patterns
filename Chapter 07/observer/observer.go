package observer

import "fmt"

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.ObserversList = append(s.ObserversList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range s.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}

func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
