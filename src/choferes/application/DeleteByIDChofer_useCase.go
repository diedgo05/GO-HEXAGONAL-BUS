package application

import "bus-project/src/choferes/domain"

type DeleteByIDChoferUseCase struct {
	db domain.IChoferesRepository
}

func NewDeleteByIDChoferUseCase(db domain.IChoferesRepository) *DeleteByIDChoferUseCase {
	return &DeleteByIDChoferUseCase{db: db}
}

func (uc *DeleteByIDChoferUseCase) Run(id int) error {
	return uc.db.DeleteByID(id)
}