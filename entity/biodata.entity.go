package entity

type Biodata struct {
	Id        int    `json:"ID"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}
