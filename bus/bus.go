package bus

// Message sent in the bus
type Message struct {
}

var bus chan Message

// Init start the vent bus
func Init() {
	bus = make(chan Message)
}

// Stop event bus
func Stop() {
	close(bus)
}

// Publish on topic
func Publish(topic string, msg Message) {
}

// Subscribe on topic
func Subscribe(topic string) chan<- Message {
	return nil
}
