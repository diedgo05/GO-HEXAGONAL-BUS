package application

import "bus-project/src/choferes/domain"

type GetAllChoferesUseCase struct {
	db domain.IChoferesRepository
}

func NewGetAllChoferesUseCase(db domain.IChoferesRepository) *GetAllChoferesUseCase {
	return &GetAllChoferesUseCase{db: db}
}

func (uc *GetAllChoferesUseCase) Run() ([]domain.Chofer, error) {
	return uc.db.FindAll()
}