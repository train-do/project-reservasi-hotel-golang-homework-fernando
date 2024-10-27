package model

type Pelanggan struct {
	Nama string `json:"nama"`
	NoHp string `json:"noHp"`
	VIP  bool   `json:"vip"`
}
type KamarHotel struct {
	NoKamar string `json:"noKamar"`
	Price   int    `json:"price"`
	Type    string `json:"type"`
}
type Reservasi struct {
	Id               int    `json:"id"`
	Nama             string `json:"nama"`
	NoHp             string `json:"noHp"`
	NoKamar          string `json:"noKamar"`
	TglReservasi     string `json:"tglReservasi"`
	TotalPembayaran  int    `json:"totalPembayaran"`
	StatusPembayaran bool   `json:"statusPembayaran"`
	StatusReservasi  bool   `json:"statusReservasi"`
	CreatedAt        string `json:"createdAt"`
}

var ListPelanggan *[]Pelanggan = &[]Pelanggan{}
var ListKamar *[]KamarHotel = &[]KamarHotel{}
var ListReservasi *[]Reservasi = &[]Reservasi{}