package application

import "bus-project/src/buses/domain"

type FindBusByIdChoferUseCase struct {
	db domain.IBusesRepository
}

func NewFindBusByIdChoferUseCase(db domain.IBusesRepository) *FindBusByIdChoferUseCase {
	return &FindBusByIdChoferUseCase{db: db}
}

func (uc *FindBusByIdChoferUseCase) Run(choferID int) ([]domain.Buses, error) {
	return uc.db.FindBusByIdChofer(choferID)
}
