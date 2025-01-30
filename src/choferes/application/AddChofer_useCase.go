package application

import "bus-project/src/choferes/domain"

type AddChoferUseCase struct {
	db domain.IChoferesRepository
}

func NewAddChoferUseCase(db domain.IChoferesRepository) *AddChoferUseCase {
	return &AddChoferUseCase{db: db}
}

func (uc *AddChoferUseCase) Run(chofer domain.Chofer) error {
	return uc.db.Save(chofer)
}
