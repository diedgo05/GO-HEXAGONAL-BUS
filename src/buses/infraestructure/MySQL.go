package infraestructure

import (
	"bus-project/src/buses/domain"
	"database/sql"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(bus domain.Buses) error {
	query := "INSERT INTO buses (idBus, placa, capacidad, choferID) VALUES (?, ?, ?, ?)"
	_, err := mysql.db.Exec(query, bus.GetID(), bus.GetPlaca(), bus.GetCapacidad(), bus.GetChoferID())

	if err != nil {
		return err
	}

	fmt.Println("Bus guardado correctamente")
	return nil
}

func (mysql *MySQL) FindAllBuses() ([]domain.Buses, error) {
	query := "SELECT idBus, placa, capacidad, choferID FROM buses"
	rows, err := mysql.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var buses []domain.Buses
	for rows.Next() {
		var bus domain.Buses
		err := rows.Scan(&bus.IdBus, &bus.Placa, &bus.Capacidad, &bus.ChoferID)
		if err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}
	fmt.Println("Choferes encontrados correctamente")
	return buses, nil
}

func (mysql *MySQL) FindBusByIdChofer(choferID int) ([]domain.Buses, error) {
	query := "SELECT idBus, placa, capacidad, choferID FROM buses WHERE choferID = ?"
	rows, err := mysql.db.Query(query, choferID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var buses []domain.Buses
	for rows.Next() {
		var bus domain.Buses
		err := rows.Scan(&bus.IdBus, &bus.Placa, &bus.Capacidad, &bus.ChoferID)
		if err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}
	fmt.Println("Buses encontrados por chofer correctamente")
	return buses, nil
}

func (mysql *MySQL) UpdateByID(idBus int, bus domain.Buses) error {
	query := "UPDATE buses SET placa = ?, capacidad = ?, choferID = ? WHERE idBus = ?"
	_, err := mysql.db.Exec(query, bus.GetPlaca(), bus.GetCapacidad(), bus.GetChoferID(), idBus)

	if err != nil {
		return err
	}

	fmt.Println("Datos del bus actualizados correctamente")
	return nil
}

func (mysql *MySQL) DeleteByID(idBus int) error {
	query := "DELETE FROM buses WHERE idBus = ?"
	_, err := mysql.db.Exec(query, idBus)

	if err != nil {
		return err
	}

	fmt.Println("Bus eliminado correctamente")
	return nil
}