package infraestructure

import (
	"bus-project/src/choferes/domain"
	"database/sql"
	"fmt"
)


type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(chofer domain.Chofer) error {
	query := "INSERT INTO choferes (nombre, apellido_p, apellido_m, edad) VALUES (?, ?, ?, ?)"
	_, err := mysql.db.Exec(query, chofer.GetNombre(), chofer.GetApellidoP(), chofer.GetApellidoM(), chofer.GetEdad())

	if err != nil {
		return err
	}

	fmt.Println("Chofer guardado correctamente")
	return nil
}

func (mysql *MySQL) FindAll() ([]domain.Chofer, error) {
	query := "SELECT * FROM choferes"
	rows, err := mysql.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var choferes []domain.Chofer
	for rows.Next() {
		var chofer domain.Chofer
		err := rows.Scan(&chofer.ID, &chofer.Nombre, &chofer.Apellido_p, &chofer.Apellido_m, &chofer.Edad)
		if err != nil {
			return nil, err
		}
		choferes = append(choferes, chofer)
	}
	fmt.Println("Choferes encontrados correctamente")
	return choferes, nil
}

func (mysql *MySQL) UpdateByID(id int, chofer domain.Chofer) error {
	query := "UPDATE choferes SET nombre = ?, apellido_p = ?, apellido_m = ?, edad = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, chofer.GetNombre(), chofer.GetApellidoP(), chofer.GetApellidoM(), chofer.GetEdad(), id)

	if err != nil {
		return err
	}

	fmt.Println("Datos del chofer actualizados correctamente")
	return nil
}

func (mysql *MySQL) DeleteByID(id int) error {
	query := "DELETE FROM choferes WHERE id = ?"
	_, err := mysql.db.Exec(query, id)

	if err != nil {
		return err
	}

	fmt.Println("Chofer eliminado correctamente")
	return nil
}