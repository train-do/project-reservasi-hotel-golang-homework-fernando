package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/model"
	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/view"
)
func ValidasiInputTanggal(y string, m string, d string, noKamar string) error {
	year, _ := strconv.Atoi(y)
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)
	m = strconv.Itoa(month)
	d = strconv.Itoa(day)
	input := LayoutDate(y, m, d)
    timestamp := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	pembanding := TimeToString(timestamp)
	if input !=  pembanding{
		return errors.New(d+"/"+m+"/"+y+ " tidak tersedia di kalender")
	}
	if model.ListReservasi != nil {
		for _, v := range *model.ListReservasi {
			if input == v.TglReservasi && v.NoKamar == noKamar{
				return errors.New("kamar sudah di booking")
			}
		}
	}
	return nil
}
func ValidasiPelanggan(nama string, noHp string) error {
	var errReturn string
	var idx int
	for i, v := range *model.ListPelanggan {
		if v.Nama + v.NoHp == nama +noHp {
			return nil
		}else if v.NoHp == noHp {
			idx = i
			errReturn = fmt.Sprintf("info pelanggan ditemukan didatabase, namun nama berbeda\nnama yang terdaftar adalah %s", v.Nama)
			view.PrintWarning(errReturn)
		}
	}
	if errReturn != "" {
		for{
			input := view.FormInput("1. Reservasi dengan nama sekarang dan update informasi nama di database\n2. Reservasi sesuai dengan nama di database\nPilih untuk melanjutkan reservasi : ")
			err := ValidationInput(input, `^(1|2)$`, "invalid input")
			if !IsSessionValid {
				return nil
			} else if err != nil {
				view.PrintError(err.Error())
				continue
			} else if input == "1" {
				(*model.ListPelanggan)[idx].Nama = nama
				EncodePelanggan()
				return nil
			} else if input == "2" {
				return errors.New((*model.ListPelanggan)[idx].Nama)
			}
			break
		}
	}
	newPelanggan := model.Pelanggan{Nama: nama, NoHp: noHp}
	*model.ListPelanggan = append(*model.ListPelanggan, newPelanggan)
	EncodePelanggan()
	return nil
}
func ValidasiLogin(username string, password string) error {
	if username != "admin" {
		return errors.New("invalid username")
	} else if password != "1234" {
		return errors.New("invalid password")
	}
	return nil
}
func ValidationInput(input string, regex string, errMessage string) error {
	r := regexp.MustCompile(regex)
	if !r.MatchString(input) {
		return fmt.Errorf("%s", errMessage)
	}
	if input == "" || input == " "{
		return errors.New("tidak boleh kosong")
	}
	return nil
}