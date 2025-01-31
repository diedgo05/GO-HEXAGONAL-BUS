package application

import "bus-project/src/buses/domain"

type AddBusUseCase struct {
	db domain.IBusesRepository
}

func NewAddBusUseCase(db domain.IBusesRepository) *AddBusUseCase {
	return &AddBusUseCase{db: db}
}

func (uc *AddBusUseCase) Run(bus domain.Buses) error {
	return uc.db.Save(bus)
}