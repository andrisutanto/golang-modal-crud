package mahasiswamodel

import (
	"database/sql"

	"github.com/andrisutanto/golang-modal-crud/config"
	"github.com/andrisutanto/golang-modal-crud/entities"
)

type MahasiswaModel struct {
	db *sql.DB
}

func New() *MahasiswaModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{db: db}
}

func (m *MahasiswaModel) FindAll(mahasiswa *[]entities.Mahasiswa) error {
	rows, err := m.db.Query("select * from mahasiswa")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Mahasiswa
		rows.Scan(
			&data.Id,
			&data.NamaLengkap,
			&data.JenisKelamin,
			&data.TempatLahir,
			&data.TanggalLahir,
			&data.Alamat)

		// setiap data yang didapatkan, akan di append ke pointer mahasiswa
		*mahasiswa = append(*mahasiswa, data)
	}

	return nil
}
