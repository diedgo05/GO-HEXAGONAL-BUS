package repositories

import "bus-project/src/buses/domain"

type IEventPublisher interface {
	PublishEvent(bus domain.Buses) error
}