package wilayah

type FindKabKotaInput struct {
	ProvinsiID int `json:"provinsi_id" binding:"required"`
}

type FindKecamatanInput struct {
	KabKotaID int `json:"kabkota_id" binding:"required"`
}

type FindKelurahanInput struct {
	KecamatanID int `json:"kecamatan_id" binding:"required"`
}
