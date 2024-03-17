package main

func main() {
	//initial requirement
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)
	//future requirement
	ws := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine: ws}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
