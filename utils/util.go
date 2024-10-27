package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/train-do/project-reservasi-hotel-golang-homework-fernando/model"
)
var IsSessionValid bool
func StartSession(ctx context.Context)  {
	// fmt.Println(IsSessionValid)
	<- ctx.Done()
	IsSessionValid = false
}
func ToJSON(data interface{}) string {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}
	return string(dataJSON)
}
func UpdateVIP(nama string, noHp string)  {
	i := 1
	for _, v := range *model.ListReservasi {
		// fmt.Println(v.Nama, noHp, "+++++++")
		if v.Nama+v.NoHp == nama+noHp && v.StatusPembayaran {
			// fmt.Println(v.Nama, noHp, v.Id, "+++++++")
			i++
		}
	}
	for idx, v := range *model.ListPelanggan {
		// fmt.Println(v.Nama, noHp, "------", i)
		if nama+noHp == v.Nama+v.NoHp && i >= 5 {
			// fmt.Println(v.Nama, noHp, "-----", i)
			(*model.ListPelanggan)[idx].VIP = true
			// fmt.Printf("%+v", model.ListPelanggan)
			// fmt.Printf("%+v", *model.ListPelanggan)
			EncodePelanggan()
		}
	}
	// fmt.Printf("%+v", model.ListPelanggan)
}
func GenerateListIdPelanggan() string {
	output := "^("
	for i, v := range *model.ListReservasi {
		if i == len(*model.ListReservasi)-1 {
			id := strconv.Itoa(v.Id)
			output += id
			break
		}
		id := strconv.Itoa(v.Id)
		output += id + "|"
	}
	output += ")$"
	// println(output)
	return output
}
func GenerateListNoKamar() string {
	output := "^("
	for i, v := range *model.ListKamar {
		if i == len(*model.ListKamar)-1 {
			output += v.NoKamar
			break
		}
		output += v.NoKamar + "|"
	}
	output += ")$"
	// println(output)
	return output
}
func GenerateTotalPembayaran(noKamar string) int {
	for _, v := range *model.ListKamar {
		if noKamar == v.NoKamar {
			return v.Price
		}
	}
	return -1
}
func StringToTime(y string, m string, d string) time.Time {
	parsedTime, err := time.Parse("2006-01-02", y+m+d)
    if err != nil {
        fmt.Println("Error parsing time:", err)
    }
	return parsedTime
}
func TimeToString(data time.Time) string {
	timeString := data.Format("2006-01-02")
	return timeString
}
func LayoutDate(y string, m string, d string) string {
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)
	m = strconv.Itoa(month)
	d = strconv.Itoa(day)
	str := y+"-"+m+"-"+d
	if month < 10 && day < 10{
		str = y+"-0"+m+"-0"+d
	}else if month < 10 {
		str = y+"-0"+m+"-"+d
	}else if day < 10{
		str = y+"-"+m+"-0"+d
	}
	return str
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}