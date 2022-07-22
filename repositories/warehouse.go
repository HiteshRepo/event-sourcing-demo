package repositories

import (
	"github.com/hiteshpattanayak-tw/EventSourcing/event"
	"github.com/hiteshpattanayak-tw/EventSourcing/model"
)

type WarehouseRepository struct {
	inMemoryStreams map[int][]event.Event
}

func NewWarehouseRepository() *WarehouseRepository {
	return &WarehouseRepository{inMemoryStreams: make(map[int][]event.Event)}
}

func (wr *WarehouseRepository) Get(id int) *model.Product {
	product := model.NewProduct(id)

	if events, ok := wr.inMemoryStreams[id]; ok {
		for _,e := range events {
			product.AddEvent(e)
		}
	}

	return product
}

func (wr *WarehouseRepository) Save(product *model.Product)  {
	wr.inMemoryStreams[product.Id] = product.GetEvents()
}
