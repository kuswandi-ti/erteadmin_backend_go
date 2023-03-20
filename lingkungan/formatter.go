package lingkungan

type LingkunganFormatter struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	ProvinsiID  int    `json:"provinsi_id"`
	KabkotaID   int    `json:"kabkota_id"`
	KecamatanID int    `json:"kecamatan_id"`
	KelurahanID int    `json:"kelurahan_id"`
}

func FormatLingkungan(lingkungan Lingkungan) LingkunganFormatter {
	formatter := LingkunganFormatter{
		ID:          lingkungan.ID,
		Nama:        lingkungan.Nama,
		Alamat:      lingkungan.Alamat,
		ProvinsiID:  lingkungan.ProvinsiID,
		KabkotaID:   lingkungan.KabkotaID,
		KecamatanID: lingkungan.KecamatanID,
		KelurahanID: lingkungan.KelurahanID,
	}

	return formatter
}
