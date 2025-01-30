package domain

type Chofer struct {
	ID        int
	Nombre    string
	Apellido_p string
	Apellido_m string
	Edad	  int
}

func NewChofer(id int, nombre string, apellido_p string, apellido_m string, edad int) *Chofer {
	return &Chofer{
		ID:       id,
		Nombre:   nombre,
		Apellido_p: apellido_p,
		Apellido_m: apellido_m,
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

func (c *Chofer) GetApellidoP() string {
	return c.Apellido_p
}
func (c *Chofer) SetApellidoP(apellido_p string) {
	c.Apellido_p = apellido_p
}

func (c *Chofer) GetApellidoM() string {
	return c.Apellido_m
}
func (c *Chofer) SetApellidoM(apellido_m string) {
	c.Apellido_m = apellido_m
}

func (c *Chofer) GetEdad() int {
	return c.Edad
}
func (c *Chofer) SetEdad(edad int) {
	c.Edad = edad
}