package application

import "bus-project/src/choferes/domain"

type GetChoferByIDUseCase struct {
	db domain.IChoferesRepository
}

func NewGetChoferByIDUseCase(db domain.IChoferesRepository) *GetChoferByIDUseCase {
	return &GetChoferByIDUseCase{db: db}
}

func (uc *GetChoferByIDUseCase) Run(id int) ([]domain.Chofer, error) {
	return uc.db.FindByID(id)
}