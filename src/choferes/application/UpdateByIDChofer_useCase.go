package application

import "bus-project/src/choferes/domain"

type UpdateByIDChoferUseCase struct {
	db domain.IChoferesRepository
}

func NewUpdateByIDChoferUseCase(db domain.IChoferesRepository) *UpdateByIDChoferUseCase {
	return &UpdateByIDChoferUseCase{db: db}
}

func (uc *UpdateByIDChoferUseCase) Run(id int, chofer domain.Chofer) error {
	return uc.db.UpdateByID(id, chofer)
}