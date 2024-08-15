package main

import (
	biodataRepo "assigment1/repository"
	"fmt"
	"os"
)

func main() {
	bio := biodataRepo.DataBiodata(os.Args[1])
	if bio.Nama != "" {
		fmt.Println("ID :", bio.Id)
		fmt.Println("nama :", bio.Nama)
		fmt.Println("alamat :", bio.Alamat)
		fmt.Println("pekerjaan :", bio.Pekerjaan)
		fmt.Println("alasan :", bio.Alasan)
	} else {
		fmt.Println("Data Not Found")
	}
}
