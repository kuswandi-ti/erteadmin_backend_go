package wilayah

import "gorm.io/gorm"

type Repository interface {
	FindAllProvinsi() ([]Provinsi, error)
	FindKabKotaByProvinsiID(ID int) ([]Kabkota, error)
	FindKecamatanByKabKotaID(ID int) ([]Kecamatan, error)
	FindKelurahanByKecamatanID(ID int) ([]Kelurahan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllProvinsi() ([]Provinsi, error) {
	var provinsis []Provinsi

	err := r.db.Find(&provinsis).Error
	if err != nil {
		return provinsis, err
	}

	return provinsis, nil
}

func (r *repository) FindKabKotaByProvinsiID(ID int) ([]Kabkota, error) {
	var kabkotas []Kabkota

	err := r.db.Where("provinsi_id = ?", ID).Find(&kabkotas).Error
	if err != nil {
		return kabkotas, err
	}

	return kabkotas, nil
}

func (r *repository) FindKecamatanByKabKotaID(ID int) ([]Kecamatan, error) {
	var kecamatans []Kecamatan

	err := r.db.Where("kabkota_id = ?", ID).Find(&kecamatans).Error
	if err != nil {
		return kecamatans, err
	}

	return kecamatans, nil
}

func (r *repository) FindKelurahanByKecamatanID(ID int) ([]Kelurahan, error) {
	var kelurahans []Kelurahan

	err := r.db.Where("kecamatan_id = ?", ID).Find(&kelurahans).Error
	if err != nil {
		return kelurahans, err
	}

	return kelurahans, nil
}
