package application

import "bus-project/src/buses/domain"

type GetAllBusesUseCase struct {
	db domain.IBusesRepository
}

func NewGetAllBusesUseCase(db domain.IBusesRepository) *GetAllBusesUseCase {
	return &GetAllBusesUseCase{db: db}
}

func (uc *GetAllBusesUseCase) Run() ([]domain.Buses, error) {
	return uc.db.FindAllBuses()
}
