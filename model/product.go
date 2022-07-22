package model

import (
	"errors"
	"github.com/hiteshpattanayak-tw/EventSourcing/event"
	"time"
)

type Product struct {
	Id           int
	events       []event.Event
	currentState *event.CurrentState
}

func NewProduct(id int) *Product {
	return &Product{Id: id, currentState: &event.CurrentState{QuantityOnHand: 0}}
}

func (p *Product) ShipProduct(quantity int) error {
	if quantity > p.currentState.QuantityOnHand {
		return errors.New("not enough products to ship")
	}

	p.AddEvent(event.NewProductShippedEvent(p.Id, quantity, time.Now().UnixMilli()))

	return nil
}

func (p *Product) ReceiveProduct(quantity int) {
	p.AddEvent(event.NewProductReceivedEvent(p.Id, quantity, time.Now().UnixMilli()))
}

func (p *Product) AdjustInventory(quantity int, reason string) error {
	if p.currentState.QuantityOnHand+quantity < 0 {
		return errors.New("cannot adjust to current quantity")
	}

	p.AddEvent(event.NewInventoryAdjustedEvent(p.Id, quantity, reason, time.Now().UnixMilli()))

	return nil
}

func (p *Product) GetCurrentState() *event.CurrentState {
	return p.currentState
}

func (p *Product) AddEvent(event event.Event) {
	event.Apply(p.currentState)
	p.events = append(p.events, event)
}

func (p *Product) GetEvents() []event.Event {
	return p.events
}
