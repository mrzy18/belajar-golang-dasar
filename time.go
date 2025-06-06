package main

import (
	"fmt"
	"time"
)

func main() {
	var timeNow = time.Now()

	// Mengembalikan informasi waktu saat ini
	fmt.Println(timeNow)
	// mengembalikan tahun dan bulan saat ini
	fmt.Println("Year: ", timeNow.Year(), "Month: ", timeNow.Month())

	// Parsing string ke time.Time
	var layoutFormat string = "02/01/2006 MST"
	var value string = "06/06/2025 WIB"
	var date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

	// Parsing time.Time ke string
	fmt.Println("time.Time to string: ", date.Format(time.RFC3339))

	// time.Sleep() untuk menghentikan program sejenak
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 seconds")
}
