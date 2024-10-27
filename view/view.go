package view

import "fmt"
func DashboardView() {
	fmt.Println("1. Show Data Pelanggan")
	fmt.Println("2. Show Room")
	fmt.Println("3. Create Reservasi")
	fmt.Println("4. Edit Reservasi")
	fmt.Println("5. Delete Reservasi")
	fmt.Println("0. Logout")
}
func FormInput(label string) string {
	var input string
	fmt.Print(label)
	fmt.Scanln(&input)
	return input
}
func PrintSucces(text string) {
	fmt.Printf("\033[1;32m %s \033[0m\n", text)
}
func PrintWarning(text string) {
	fmt.Printf("\033[33m %s \033[0m\n", text)
}
func PrintError(text string) {
	fmt.Printf("\033[31m %s \033[0m\n", text)
}