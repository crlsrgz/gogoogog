package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	john := Passenger{"John", 1, false}
	fmt.Println(john)

	var bill = Passenger{Name: "Bill", TicketNumber: 2}
	var jane = Passenger{Name: "Jane", TicketNumber: 3}

	fmt.Println(bill, jane)

	var jamie Passenger
	jamie.Name = "Jamie"
	jamie.TicketNumber = 4

	fmt.Println(jamie)

	john.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill boarded")

	}

	if john.Boarded {
		fmt.Println("john boarded")

	}

	jamie.Boarded = true
	bus := Bus{jamie}
	//Access the Passager from the bus struct
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

}
