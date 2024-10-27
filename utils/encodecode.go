package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/model"
)

func EncodePelanggan() {
	file, err := os.Create("pelanggan.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(model.ListPelanggan); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
func EncodeReservasi() {
	file, err := os.Create("reservasi.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(model.ListReservasi); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
func DecodePelanggan() {
	file, err := os.Open("pelanggan.json")
	if err != nil {
		// fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&model.ListPelanggan); err != nil {
		// fmt.Println("Error decoding JSON:", err)
		return
	}
}
func DecodeKamar() {
	file, err := os.Open("kamar.json")
	if err != nil {
		// fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&model.ListKamar); err != nil {
		// fmt.Println("Error decoding JSON:", err)
		return
	}
}
func DecodeReservasi() {
	file, err := os.Open("reservasi.json")
	if err != nil {
		// fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&model.ListReservasi); err != nil {
		// fmt.Println("Error decoding JSON:", err)
		return
	}
}