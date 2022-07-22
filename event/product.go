package event

import (
	"log"
)

type productShippedEvent struct {
	id           int
	quantity     int
	datetime     int64
}

func NewProductShippedEvent(id int, quantity int, datetime int64) Event {
	return productShippedEvent{id: id, quantity: quantity, datetime: datetime}
}

func (pse productShippedEvent) Apply(currState *CurrentState) {
	currState.QuantityOnHand -= pse.quantity
}

func (pse productShippedEvent) Display() {
	log.Printf("Product (%d) shipped with quantity: %d at %d\n", pse.id, pse.quantity, pse.datetime)
}

type productReceivedEvent struct {
	id           int
	quantity     int
	datetime     int64
}

func NewProductReceivedEvent(id int, quantity int, datetime int64) Event {
	return productReceivedEvent{id: id, quantity: quantity, datetime: datetime}
}

func (pre productReceivedEvent) Apply(currState *CurrentState) {
	currState.QuantityOnHand += pre.quantity
}

func (pre productReceivedEvent) Display() {
	log.Printf("Product (%d) received with quantity: %d at %d\n", pre.id, pre.quantity, pre.datetime)
}

type inventoryAdjustedEvent struct {
	id           int
	quantity     int
	datetime     int64
	reason       string
}

func NewInventoryAdjustedEvent(id int, quantity int, reason string, datetime int64) Event {
	return inventoryAdjustedEvent{id: id, quantity: quantity, reason: reason, datetime: datetime}
}

func (iae inventoryAdjustedEvent) Apply(currState *CurrentState) {
	currState.QuantityOnHand += iae.quantity
}

func (iae inventoryAdjustedEvent) Display() {
	log.Printf("Product (%d) adjusted with quantity: %d at %d\n because of %s", iae.id, iae.quantity, iae.datetime, iae.reason)
}
