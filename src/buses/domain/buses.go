package domain

type Buses struct {
	IdBus        int
	Placa       string
	Capacidad   int
	ChoferID    int
}

func NewBus(idBus int, placa string, capacidad int, choferID int) *Buses {
	return &Buses{
		IdBus:      idBus,
		Placa:      placa,
		Capacidad:  capacidad,
		ChoferID:   choferID,
	}
}

func (b *Buses) GetID() int {
	return b.IdBus
}
func (b *Buses) SetID(idBus int) {
	b.IdBus = idBus
}

func (b *Buses) GetPlaca() string {	
	return b.Placa
}
func (b *Buses) SetPlaca(placa string) {
	b.Placa = placa
}

func (b *Buses) GetCapacidad() int {
	return b.Capacidad
}
func (b *Buses) SetCapacidad(capacidad int) {
	b.Capacidad = capacidad
}

func (b *Buses) GetChoferID() int {
	return b.ChoferID
}
func (b *Buses) SetChoferID(choferID int) {
	b.ChoferID = choferID
}