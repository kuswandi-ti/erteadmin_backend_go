package lingkungan

import "errors"

type Service interface {
	SaveLingkungan(input CreateLingkunganInput) (Lingkungan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveLingkungan(input CreateLingkunganInput) (Lingkungan, error) {
	lingkungan := Lingkungan{}

	lingkungan.Nama = input.Nama
	lingkungan.Alamat = input.Alamat
	lingkungan.ProvinsiID = input.ProvinsiID
	lingkungan.KabkotaID = input.KabkotaID
	lingkungan.KecamatanID = input.KecamatanID
	lingkungan.KelurahanID = input.KelurahanID

	lingkungan, err := s.repository.FindByNama(input.Nama)
	if err != nil {
		return lingkungan, err
	}
	if lingkungan.ID != 0 {
		return lingkungan, errors.New("Nama lingkungan sudah terdaftar")
	}

	newLingkungan, err := s.repository.Save(lingkungan)
	if err != nil {
		return newLingkungan, err
	}

	return newLingkungan, nil
}
