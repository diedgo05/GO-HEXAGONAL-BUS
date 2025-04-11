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

func (mysql *MySQL) Save(bus domain.Buses) (int, error) {
	query := "INSERT INTO buses ( placa, capacidad, disponible, choferID) VALUES (?, ?, ?, ?)"
	result, err := mysql.db.Exec(query, bus.Placa, bus.Capacidad, bus.Disponible, bus.ChoferID)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (mysql *MySQL) FindAllBuses() ([]domain.Buses, error) {
	query := "SELECT idBus, placa, capacidad, disponible, choferID FROM buses"

	rows, err := mysql.db.Query(query)
	fmt.Println("Ejecutando query:", query)
	
	if err != nil {
		fmt.Println("Error al ejecutar el query:", err)
		return nil, err
	}
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var buses []domain.Buses
	for rows.Next() {
		var bus domain.Buses
		err := rows.Scan(&bus.IdBus, &bus.Placa, &bus.Capacidad,&bus.Disponible , &bus.ChoferID)
		if err != nil {
			fmt.Println("Error al escanear fila:", err)
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
	query := "UPDATE buses SET  disponible =? WHERE idBus = ?"
	_, err := mysql.db.Exec(query,  bus.Disponible, idBus)

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