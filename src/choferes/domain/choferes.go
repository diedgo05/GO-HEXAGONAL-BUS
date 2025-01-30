package domain

type Chofer struct {
	ID        int
	Nombre    string
	Apellido  string
	Edad	  int
}

func NewChofer(id int, nombre string, apellido string, edad int) *Chofer {
	return &Chofer{
		ID:       1,
		Nombre:   nombre,
		Apellido: apellido,
		Edad:     edad,
	}
}

func (c *Chofer) GetID() int {
	return c.ID
}
func (c *Chofer) SetID(id int) {
	c.ID = id
}

func (c *Chofer) GetNombre() string {
	return c.Nombre
}
func (c *Chofer) SetNombre(nombre string) {
	c.Nombre = nombre
}

func (c *Chofer) GetApellido() string {
	return c.Apellido
}
func (c *Chofer) SetApellido(apellido string) {
	c.Apellido = apellido
}

func (c *Chofer) GetEdad() int {
	return c.Edad
}
func (c *Chofer) SetEdad(edad int) {
	c.Edad = edad
}