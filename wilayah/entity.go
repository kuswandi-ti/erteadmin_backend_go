package wilayah

import "time"

type Provinsi struct {
	ID        int
	Nama      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Kabkota struct {
	ID         int
	Nama       string
	ProvinsiID int
	Provinsi   Provinsi
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Kecamatan struct {
	ID        int
	Nama      string
	KabkotaID int
	Kabkota   Kabkota
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Kelurahan struct {
	ID          int
	Nama        string
	KecamatanID int
	Kecamatan   Kecamatan
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
