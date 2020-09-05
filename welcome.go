package main

import (
	"fmt"
	"time"
)

func main() {

	var ad string
	// for english "enter your name"
	fmt.Print("Adınızı Giriniz : ")
	fmt.Scanln(&ad)
	s := time.Now()
	switch {
	case s.Hour() < 11:
		// for english "good morning"
		fmt.Println("Günaydın", ad)
	case s.Hour() < 16:
		// for english "good afternoon"
		fmt.Println("Tünaydın", ad)
	case s.Hour() > 16:
		// for english "good evening"
		fmt.Println("İyi Akşamlar", ad)
	}

}
