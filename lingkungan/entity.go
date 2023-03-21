package lingkungan

import (
	"erteadmin_backend/wilayah"
	"time"
)

type Lingkungan struct {
	ID          int
	Nama        string
	Alamat      string
	ProvinsiID  int
	KabkotaID   int
	KecamatanID int
	KelurahanID int
	Provinsi    wilayah.Provinsi
	Kabkota     wilayah.Kabkota
	Kecamatan   wilayah.Kecamatan
	Kelurahan   wilayah.Kelurahan
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
