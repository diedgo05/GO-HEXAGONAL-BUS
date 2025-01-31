package application

import "bus-project/src/buses/domain"

type DeleteBusByIDUseCase struct {
	db domain.IBusesRepository
}

func NewDeleteBusByIDUseCase(db domain.IBusesRepository) *DeleteBusByIDUseCase {
	return &DeleteBusByIDUseCase{db: db}
}

func (uc *DeleteBusByIDUseCase) Run(idBus int) error {
	return uc.db.DeleteByID(idBus)
}