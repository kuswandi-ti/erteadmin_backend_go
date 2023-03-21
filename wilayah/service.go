package wilayah

type Service interface {
	GetAllProvinsi() ([]Provinsi, error)
	// GetKabKotaByProvinsiID(ID int) ([]Kabkota, error)
	GetKabKotaByProvinsiID(input FindKabKotaInput) ([]Kabkota, error)
	// GetKecamatanByKabKotaID(ID int) ([]Kecamatan, error)
	GetKecamatanByKabKotaID(FindKecamatanInput) ([]Kecamatan, error)
	GetKelurahanByKecamatanID(FindKelurahanInput) ([]Kelurahan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllProvinsi() ([]Provinsi, error) {
	provinsis, err := s.repository.FindAllProvinsi()
	if err != nil {
		return provinsis, err
	}

	return provinsis, nil
}

// func (s *service) GetKabKotaByProvinsiID(ID int) ([]Kabkota, error) {
// 	kabkotas, err := s.repository.FindKabKotaByProvinsiID(ID)
// 	if err != nil {
// 		return kabkotas, err
// 	}

// 	return kabkotas, nil
// }

func (s *service) GetKabKotaByProvinsiID(input FindKabKotaInput) ([]Kabkota, error) {
	provinsi_id := input.ProvinsiID

	kabkotas, err := s.repository.FindKabKotaByProvinsiID(provinsi_id)
	if err != nil {
		return kabkotas, err
	}

	return kabkotas, nil
}

// func (s *service) GetKecamatanByKabKotaID(ID int) ([]Kecamatan, error) {
// 	kecamatans, err := s.repository.FindKecamatanByKabKotaID(ID)
// 	if err != nil {
// 		return kecamatans, err
// 	}

// 	return kecamatans, nil
// }

func (s *service) GetKecamatanByKabKotaID(input FindKecamatanInput) ([]Kecamatan, error) {
	kabkota_id := input.KabKotaID

	kecamatans, err := s.repository.FindKecamatanByKabKotaID(kabkota_id)
	if err != nil {
		return kecamatans, err
	}

	return kecamatans, nil
}

func (s *service) GetKelurahanByKecamatanID(input FindKelurahanInput) ([]Kelurahan, error) {
	kecamatan_id := input.KecamatanID

	kelurahans, err := s.repository.FindKelurahanByKecamatanID(kecamatan_id)
	if err != nil {
		return kelurahans, err
	}

	return kelurahans, nil
}
