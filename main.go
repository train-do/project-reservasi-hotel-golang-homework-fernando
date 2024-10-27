package main

import (
	"context"
	"time"

	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/service"
	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/utils"
	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/view"
)

func main() {
	for {
		input := view.FormInput("Masukkan '0' untuk exit\nPress any key to continue...\n")
		if input == "0" {
			break
		} else {
			for{
				username := view.FormInput("Username : ")
				password := view.FormInput("Password : ")
				err := utils.ValidasiLogin(username, password)
				if err != nil {
					view.PrintError(err.Error())
					continue
				}
				utils.ClearScreen()
				break
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			utils.IsSessionValid = true
			go utils.StartSession(ctx)
			service.Dashboard(cancel)
		}
	}
}
func init()  {
	utils.DecodePelanggan()
	utils.DecodeKamar()
	utils.DecodeReservasi()
	utils.ClearScreen()
}
