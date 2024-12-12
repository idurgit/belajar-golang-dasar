package main

import "fmt"

func main() {
	names := [...]string{"Jaka", "Kelana", "Umbara", "Eko", "Kurniawan", "Khannedy"}

	slice1 := names[4:6]
	fmt.Println(slice1) 

	slice2 := names[:3] 
	fmt.Println(slice2) 

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]   
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Senin Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1) 
	fmt.Println(daySlice2)

	fmt.Println(days)

}
