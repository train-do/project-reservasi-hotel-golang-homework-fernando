package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/model"
	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/utils"
	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/view"
)

func Dashboard(cancel context.CancelFunc) {
	for{
		if utils.IsSessionValid {
			view.DashboardView()
			input := view.FormInput("Pilih menu : ")
			err := utils.ValidationInput(input, "^[012345]$", "opsi menu tidak tersedia")
			if !utils.IsSessionValid {
				utils.ClearScreen()
				fmt.Println("Session berakhir, silahkan Login kembali")
				return
			} else if err != nil {
				view.PrintError(err.Error())
				continue
			} else if input == "1" {
				utils.ClearScreen()
				listPelanggan := utils.ToJSON(model.ListPelanggan)
				fmt.Println(listPelanggan)
			} else if input == "2" {
				utils.ClearScreen()
				listPelanggan := utils.ToJSON(model.ListKamar)
				fmt.Println(listPelanggan)
			} else if input == "3" {
				utils.ClearScreen()
				createReservasi()
			} else if input == "4" {
				utils.ClearScreen()
				editReservasi()
			} else if input == "5" {
				utils.ClearScreen()
				deleteReservasi()
			} else if input == "0" {
				cancel()
				utils.ClearScreen()
				fmt.Println("Berhasil Logout")
				return
			} 
		}else{
			utils.ClearScreen()
			fmt.Println("Session berakhir, silahkan Login kembali")
			break
		}
	}
}
func createReservasi(){
	var nama, noHp, noKamar, tgl, bln, thn string
	var statBayar bool
	for{
		nama = view.FormInput("Nama : ")
		err := utils.ValidationInput(nama, `^([A-Za-z]*\s?[A-Za-z]*$|^[A-Za-z]+$|^\S)$`, "invalid input")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		}
		break
	}
	for{
		noHp = view.FormInput("No. Handphone : ")
		err := utils.ValidationInput(noHp, `^\d{12}$`, "invalid input dan harus 12 angka")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		}
		break
	}
	for{
		noKamar = view.FormInput("No. Kamar : ")
		err := utils.ValidationInput(noKamar, utils.GenerateListNoKamar(), "invalid input")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		}
		break
	}
	totalPembayaran := utils.GenerateTotalPembayaran(noKamar)
	for{
		tahun:
		for{
			thn = view.FormInput("Tahun : ")
			err := utils.ValidationInput(thn, `^\d{4}$`, "invalid input")
			if !utils.IsSessionValid {
				return
			} else if err != nil {
				view.PrintError(err.Error())
				continue tahun
			}
			break
		}
		bulan:
		for{
			bln = view.FormInput("Bulan : ")
			err := utils.ValidationInput(bln, `^(0?[1-9]|1[0-2])$`, "invalid input")
			if !utils.IsSessionValid {
				return
			} else if err != nil {
				view.PrintError(err.Error())
				continue bulan
			}
			break
		}
		tanggal:
		for{
			tgl = view.FormInput("Tanggal : ")
			err := utils.ValidationInput(tgl, `^(0?[1-9]|[12][0-9]|3[01])$`, "invalid input")
			if !utils.IsSessionValid {
				return
			} else if err != nil {
				view.PrintError(err.Error())
				continue tanggal
			}
			break
		}
		err := utils.ValidasiInputTanggal(thn, bln, tgl, noKamar)
		if err != nil {
			view.PrintError(err.Error())
			continue
		}
		break
	}
	for{
		isPay := view.FormInput(fmt.Sprintf("Total Pembayaran %d rupiah\nApakah ingin melakukan pembayaran? (y/n)", totalPembayaran))
		err := utils.ValidationInput(isPay, `^(y|n)$`, "invalid input")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		} else if isPay == "y" {
			statBayar = true
			go utils.UpdateVIP(nama, noHp)
		} else if isPay == "n" {
			statBayar = false
		}
		break
	}
	err := utils.ValidasiPelanggan(nama, noHp)
	if err != nil {
		nama = err.Error()
	}
	newReservasi := model.Reservasi{
		Id: len(*model.ListReservasi),
		Nama: nama,
		NoHp: noHp,
		NoKamar: noKamar,
		TglReservasi: utils.LayoutDate(thn, bln, tgl),
		TotalPembayaran: totalPembayaran,
		StatusPembayaran: statBayar,
		StatusReservasi: true,
		CreatedAt: utils.TimeToString(time.Now()),
	}
	*model.ListReservasi = append(*model.ListReservasi, newReservasi)
	utils.EncodeReservasi()
	utils.ClearScreen()
	view.PrintSucces("Reservasi Succes")
}
func editReservasi()  {
	for{
		listReservasi := utils.ToJSON(model.ListReservasi)
		input := view.FormInput(listReservasi+"\nMasukkan Id Reservasi untuk melakukan pembayaran : ")
		err := utils.ValidationInput(input, utils.GenerateListIdPelanggan(), "invalid input")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		}
		idx, _ := strconv.Atoi(input)
		if (*model.ListReservasi)[idx].StatusPembayaran {
			utils.ClearScreen()
			view.PrintSucces("Reservasi sudah dibayarkan, tidak perlu bayar lagi")
			return
		}else{
			for{
				isPay := view.FormInput(fmt.Sprintf("Total Pembayaran %d rupiah\nApakah ingin melakukan pembayaran? (y/n)", (*model.ListReservasi)[idx].TotalPembayaran))
				err := utils.ValidationInput(isPay, `^(y|n)$`, "invalid input")
				if !utils.IsSessionValid {
					return
				} else if err != nil {
					view.PrintError(err.Error())
					continue
				} else if isPay == "y" {
					(*model.ListReservasi)[idx].StatusPembayaran = true
					utils.EncodeReservasi()
					go utils.UpdateVIP((*model.ListReservasi)[idx].Nama, (*model.ListReservasi)[idx].NoHp)
					utils.ClearScreen()
					view.PrintSucces("Pembayaran Succes")
					return
				} else if isPay == "n" {
					utils.ClearScreen()
					return
				}
			}
		}
	}
	// utils.ClearScreen()
}
func deleteReservasi()  {
	for{
		listReservasi := utils.ToJSON(model.ListReservasi)
		input := view.FormInput(listReservasi+"\nMasukkan Id Reservasi untuk melakukan pembayaran : ")
		err := utils.ValidationInput(input, utils.GenerateListIdPelanggan(), "invalid input")
		if !utils.IsSessionValid {
			return
		} else if err != nil {
			view.PrintError(err.Error())
			continue
		}
		idx, _ := strconv.Atoi(input)
		if (*model.ListReservasi)[idx].StatusPembayaran {
			utils.ClearScreen()
			view.PrintWarning("Reservasi sudah dibayarkan, tidak bisa didelete")
		}else{
			for{
				delete := view.FormInput(fmt.Sprintf("Yakin delete reservasi dengan id '%d'? (y/n)", idx))
				err := utils.ValidationInput(delete, `^(y|n)$`, "invalid input")
				if !utils.IsSessionValid {
					return
				} else if err != nil {
					view.PrintError(err.Error())
					continue
				} else if delete == "y" {
					*model.ListReservasi = append((*model.ListReservasi)[:idx], (*model.ListReservasi)[idx+1:]...)
					utils.EncodeReservasi()
					utils.ClearScreen()
					view.PrintSucces("Delete Success")
					break
				} else if delete == "n" {
					utils.ClearScreen()
					return
				}
			}
		}
		break
	}
}