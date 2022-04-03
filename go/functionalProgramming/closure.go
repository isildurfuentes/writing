package main

import (
	"fmt"
	"time"
)

func getFormattedTime() string {
	return time.Now().Format("03:04:05")
}

func checkIn(name string) func() {

	message := name + " checked in at " + getFormattedTime()

	closureFunc := func() {
		fmt.Println(message)
	}
	return closureFunc
}

func main() {

	fmt.Println("Just before checkin: ", getFormattedTime())

	mariaCheckIn := checkIn("Maria")

	time.Sleep(time.Second * 5)
	fmt.Println("After timeout: ", getFormattedTime())

	mariaCheckIn()

}
