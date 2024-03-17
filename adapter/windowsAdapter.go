package main

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	wa.windowsMachine.insertIntoUSBPort()
}
