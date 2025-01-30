package domain

type IChoferesRepository interface {
	Save(chofer Chofer) error
	FindAll() ([]Chofer, error)
	UpdateByID(id int, chofer Chofer) error
	DeleteByID(id int) error
}