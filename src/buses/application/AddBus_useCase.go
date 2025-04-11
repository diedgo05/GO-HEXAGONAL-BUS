package application

import "bus-project/src/buses/domain"

type AddBusUseCase struct {
	db domain.IBusesRepository
}

func NewAddBusUseCase(db domain.IBusesRepository) *AddBusUseCase {
	return &AddBusUseCase{db: db}
}

func (uc *AddBusUseCase) Run(bus domain.Buses) (int, error) {
	id, err := uc.db.Save(bus)
	if err != nil {
		return 0, err
	}
	return id, nil
	}