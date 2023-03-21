package lingkungan

type CreateLingkunganInput struct {
	Nama        string `json:"nama" binding:"required"`
	Alamat      string `json:"alamat" binding:"required"`
	ProvinsiID  int    `json:"provinsi_id" binding:"required"`
	KabkotaID   int    `json:"kabkota_id" binding:"required"`
	KecamatanID int    `json:"kecamatan_id" binding:"required"`
	KelurahanID int    `json:"kelurahan_id" binding:"required"`
}

type FindKabKotaInput struct {
	LingkunganID int `json:"lingkungan_id" binding:"required"`
}
