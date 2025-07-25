package main

import (
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

// EventDispatcherAdapter adapts events.EventDispatcher to work with usecase interfaces
type EventDispatcherAdapter struct {
	dispatcher *events.EventDispatcher
}

func NewEventDispatcherAdapter(dispatcher *events.EventDispatcher) *EventDispatcherAdapter {
	return &EventDispatcherAdapter{
		dispatcher: dispatcher,
	}
}

func (a *EventDispatcherAdapter) Register(eventName string, handler usecase.EventHandlerInterface) error {
	// Create an adapter for the handler
	handlerAdapter := &EventHandlerAdapter{handler: handler}
	return a.dispatcher.Register(eventName, handlerAdapter)
}

func (a *EventDispatcherAdapter) Dispatch(event usecase.OrderCreatedEvent) error {
	// Convert usecase event to events.EventInterface
	eventAdapter := &EventAdapter{event: event}
	return a.dispatcher.Dispatch(eventAdapter)
}

// EventHandlerAdapter adapts usecase.EventHandlerInterface to events.EventHandlerInterface
type EventHandlerAdapter struct {
	handler usecase.EventHandlerInterface
}

func (h *EventHandlerAdapter) Handle(event events.EventInterface) {
	// Convert events.EventInterface back to usecase.OrderCreatedEvent
	eventAdapter := &UsecaseEventAdapter{event: event}
	h.handler.Handle(eventAdapter)
}

// EventAdapter adapts usecase.OrderCreatedEvent to events.EventInterface
type EventAdapter struct {
	event usecase.OrderCreatedEvent
}

func (e *EventAdapter) GetName() string {
	return e.event.GetName()
}

func (e *EventAdapter) GetDateTime() string {
	return e.event.GetDateTime()
}

func (e *EventAdapter) GetPayload() interface{} {
	return e.event.GetPayload()
}

// UsecaseEventAdapter adapts events.EventInterface to usecase.OrderCreatedEvent
type UsecaseEventAdapter struct {
	event events.EventInterface
}

func (u *UsecaseEventAdapter) GetName() string {
	return u.event.GetName()
}

func (u *UsecaseEventAdapter) GetDateTime() string {
	return u.event.GetDateTime()
}

func (u *UsecaseEventAdapter) GetPayload() interface{} {
	return u.event.GetPayload()
}

func (u *UsecaseEventAdapter) SetPayload(payload interface{}) {
	// This is a limitation - events.EventInterface doesn't have SetPayload
	// For now, we'll ignore this
}
