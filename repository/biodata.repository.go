package repository

import (
	"assigment1/entity"
	"strings"
)

func DataBiodata(name string) (bio entity.Biodata) {
	dataBio := []entity.Biodata{
		{Id: 0, Nama: "Thomas", Alamat: "Kota A", Pekerjaan: "Programmer", Alasan: "Alasan Thomas"},
		{Id: 1, Nama: "Mbappe", Alamat: "Kediri", Pekerjaan: "Pemain Bola", Alasan: "Alasan Mbappe"},
		{Id: 2, Nama: "James", Alamat: "Malang", Pekerjaan: "Tukang Es", Alasan: "Alasan James"},
		{Id: 3, Nama: "Kiki", Alamat: "Cikampek", Pekerjaan: "Reporter", Alasan: "Alasan Kiki"},
	}

	for _, v := range dataBio {
		if strings.ToLower(name) == strings.ToLower(v.Nama) {
			return v
		}
	}
	return

}
