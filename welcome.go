package main
import (
	"fmt"
	"time"
)
func main() {
	s := time.Now()
	switch {
	case s.Hour() < 11:
		// for english "good morning"
		fmt.Println("Günaydın")
	case s.Hour() < 16:
		// for english "good afternoon"
		fmt.Println("Tünaydın")
	case s.Hour() > 16:
		// for english "good evening"
		fmt.Println("İyi Akşamlar")
	}
	time.Sleep(10 * time.Second)
}
