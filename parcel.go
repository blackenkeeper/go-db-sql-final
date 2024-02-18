package main

import (
	"database/sql"
	"log"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	insertQuery := "insert into parcel (client, status, address, created_at) values (?, ?, ?, ?)"
	res, err := s.db.Exec(insertQuery, p.Client, p.Status, p.Address, p.CreatedAt)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int(id), err
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	selectQuery := "select client, status, address, created_at from parcel where number = :number"
	row := s.db.QueryRow(selectQuery, sql.Named("number", number))

	p := Parcel{}
	err := row.Scan(&p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		log.Println(err)
		return p, err
	}

	return p, err
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	var res []Parcel

	selectQuery := "select number, client, status, address, created_at from parcel where client = :client"
	rows, err := s.db.Query(selectQuery, sql.Named("client", client))
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Parcel{}
		err = rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			log.Println(err)
			return res, err
		}

		res = append(res, p)
	}

	return res, err
}

func (s ParcelStore) SetStatus(number int, status string) error {
	updateQuery := "update parcel set status = :status where number = :number"
	_, err := s.db.Exec(updateQuery, sql.Named("status", status), sql.Named("number", number))
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (s ParcelStore) SetAddress(number int, address string) error {
	updateQuery := "update parcel set address = :address where number = :number"
	_, err := s.db.Exec(updateQuery, sql.Named("address", address), sql.Named("number", number))
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (s ParcelStore) Delete(number int) error {
	deleteQuery := "delete from parcel where number = :number and status = :status"
	_, err := s.db.Exec(deleteQuery, sql.Named("number", number), sql.Named("status", ParcelStatusRegistered))
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
