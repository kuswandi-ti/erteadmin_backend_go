package wilayah

type ProvinsiFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatProvinsiOne(provinsi Provinsi) ProvinsiFormatter {
	formatter := ProvinsiFormatter{}
	formatter.ID = provinsi.ID
	formatter.Nama = provinsi.Nama

	return formatter
}

func FormatProvinsiMany(provinsis []Provinsi) []ProvinsiFormatter {
	if len(provinsis) == 0 {
		return []ProvinsiFormatter{}
	}

	var provinsisFormatter []ProvinsiFormatter

	for _, provinsi := range provinsis {
		formatter := FormatProvinsiOne(provinsi)
		provinsisFormatter = append(provinsisFormatter, formatter)
	}

	return provinsisFormatter
}

type KabKotaFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatKabKotaOne(kabkota Kabkota) KabKotaFormatter {
	formatter := KabKotaFormatter{}
	formatter.ID = kabkota.ID
	formatter.Nama = kabkota.Nama

	return formatter
}

func FormatKabKotaMany(kabkotas []Kabkota) []KabKotaFormatter {
	if len(kabkotas) == 0 {
		return []KabKotaFormatter{}
	}

	var kabkotasFormatter []KabKotaFormatter

	for _, kabkota := range kabkotas {
		formatter := FormatKabKotaOne(kabkota)
		kabkotasFormatter = append(kabkotasFormatter, formatter)
	}

	return kabkotasFormatter
}

type KecamatanFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatKecamatanOne(kecamatan Kecamatan) KecamatanFormatter {
	formatter := KecamatanFormatter{}
	formatter.ID = kecamatan.ID
	formatter.Nama = kecamatan.Nama

	return formatter
}

func FormatKecamatanMany(kecamatans []Kecamatan) []KecamatanFormatter {
	if len(kecamatans) == 0 {
		return []KecamatanFormatter{}
	}

	var kecamatansFormatter []KecamatanFormatter

	for _, kecamatan := range kecamatans {
		formatter := FormatKecamatanOne(kecamatan)
		kecamatansFormatter = append(kecamatansFormatter, formatter)
	}

	return kecamatansFormatter
}

type KelurahanFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatKelurahanOne(kelurahan Kelurahan) KelurahanFormatter {
	formatter := KelurahanFormatter{}
	formatter.ID = kelurahan.ID
	formatter.Nama = kelurahan.Nama

	return formatter
}

func FormatKelurahanMany(kelurahans []Kelurahan) []KelurahanFormatter {
	if len(kelurahans) == 0 {
		return []KelurahanFormatter{}
	}

	var kelurahansFormatter []KelurahanFormatter

	for _, kelurahan := range kelurahans {
		formatter := FormatKelurahanOne(kelurahan)
		kelurahansFormatter = append(kelurahansFormatter, formatter)
	}

	return kelurahansFormatter
}
