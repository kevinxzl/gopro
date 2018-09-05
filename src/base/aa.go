package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
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

func Disconnected(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected!---", v.name)
	default:
		fmt.Println("Unknown Device!---")
	}
}
func main() {
	//var usb USB
	usb := PhoneConnecter{"Phone Connector"}
	usb.Connect()
	Disconnected(usb)
}

