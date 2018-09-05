package main

import "fmt"

type USB interface {
	Name() string
	Connect()
	//XXX()
}

type PhoneConnecter struct {
	name string
}

func (phone PhoneConnecter) Name() string {
	return phone.name
}

func (phone PhoneConnecter) Connect() {
	fmt.Println("Connected!---:", phone.name)
}

func Disconnected(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected!---", pc.name)
	}else{
		fmt.Println("Unknown Device!---")
	}
}
func main() {

	usb := PhoneConnecter{"Phone Connector"}
	usb.Connect()
	Disconnected(usb)
}

