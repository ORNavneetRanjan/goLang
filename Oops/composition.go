package main

import "fmt"

type Address struct {
	Street	string
	City		string
	State		string
	ZipCode int
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s, %d",
	 a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	Id		int
	Name	string
	Email	string
	BillingAddress	Address
	ShippingAddress	Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer Id: %d\n", c.Id)
	fmt.Printf("Customer Name: %s\n", c.Name)
	fmt.Printf("Customer Email: %s\n", c.Email)
	fmt.Printf("Billing Address: %s\n", c.BillingAddress.FullAddress())
	fmt.Printf("Shipping Address: %s\n", c.ShippingAddress.FullAddress())

}

func main6(){

	fmt.Println("-------- Composition ---------")
	cust1 := Customer {
		Id: 1,
		Name: "Navneet",
		Email: "abc@gmail.com",
		BillingAddress: Address{
			Street: "Anepalya",
			City: "Bangaluru",
			State: "Karnataka",
			ZipCode: 500421,
		},
		ShippingAddress: Address{
			Street: "Anepalya",
			City: "Bangaluru",
			State: "Karnataka",
			ZipCode: 500421,
		},
	}
	cust1.PrintDetails()

}