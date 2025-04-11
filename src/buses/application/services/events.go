package services

import (
	"bus-project/src/buses/application/repositories"
	"bus-project/src/buses/domain"
)

type Event struct {
	rabbit repositories.IEventPublisher
}

func NewEvent(rabbit repositories.IEventPublisher) *Event {
	return &Event{rabbit:rabbit}
}

func (e *Event) Run(bus domain.Buses) error {
return e.rabbit.PublishEvent(bus)
}