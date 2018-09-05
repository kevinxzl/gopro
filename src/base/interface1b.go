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
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected!---", pc.name)
	}else{
		fmt.Println("Unknown Device!---")
	}

}
func main() {
	phone := PhoneConnecter{"Phone Connector"}
	var con Connecter
	con = Connecter(phone)
	con.Connect()
	Disconnected(con)

	phone.name = "ABCDE"
	con.Connect()
}

