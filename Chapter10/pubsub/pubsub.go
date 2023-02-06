package pubsub

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type Publisher interface {
	start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}
