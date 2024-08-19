package events

import (
	"sync"
	"time"
)

// Evento (Carrega Dados)
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// Operações que serão executadas quando um evento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// Despachar Operações
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
