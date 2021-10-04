package events

import "sync"

type dataChannel chan DataEvent

type EventBus struct {
	subscribers map[string][]dataChannel
	sync        sync.RWMutex
}

type DataEvent struct {
	Topic string
	Data  interface{}
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]dataChannel),
	}
}

func (eb *EventBus) Subscribe(topic string, ch dataChannel) {
	eb.sync.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]dataChannel{}, ch)
	}
	eb.sync.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.sync.Lock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append([]dataChannel{}, chans...)
		go func(data DataEvent, dataChannels []dataChannel) {
			for _, ch := range dataChannels {
				ch <- data
			}
		}(DataEvent{Topic: topic, Data: data}, channels)
	}
	eb.sync.Unlock()
}
