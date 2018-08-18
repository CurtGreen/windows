package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	specialPopup := exec.Command("./windows", "500", "500", "Special Popup", "Special Popup Message...")
	fmt.Println("Initiating the popup window!")
	err := specialPopup.Run()
	if err != nil {
		fmt.Println("Something went wrong...")
	}
	fmt.Println("Hopefully we have closed the window!")
	time.Sleep(4 * time.Second)
	fmt.Println("Exiting main...")
}
