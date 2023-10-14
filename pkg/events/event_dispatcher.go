package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	Handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	/*
		First check if exists an event with the same name
		if yes, enter the method and check if the handler is already registered inside the event
		if yes, return an error
	*/
	if _, ok := ed.Handlers[eventName]; ok {
		for _, h := range ed.Handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.Handlers[eventName] = append(ed.Handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.Handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.Handlers[eventName]; ok {
		for _, h := range ed.Handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.Handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.Handlers[eventName]; ok {
		for i, h := range ed.Handlers[eventName] {
			if h == handler {
				ed.Handlers[eventName] = append(ed.Handlers[eventName][:i], ed.Handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil
}
