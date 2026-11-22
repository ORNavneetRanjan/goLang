package main

import "fmt"

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s\n Phone No: %s\n",
	 ci.Email, ci.Phone)
}

type Company struct {
	Name 					string
	Address
	ContactInfo
	BusinessType 	string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name	string
	Email string
	Address
	ContactInfo
}

func main7() {
	fmt.Println("---- Struct Embedding ----")

	com := Company{
		Name: "Rakuten",
		BusinessType: "E-commerce",
		Address: Address{
			Street: "Ring Road",
			City: "Bangalore",
			State: "KR",
			ZipCode: 500481,
		},
		ContactInfo: ContactInfo{
			Email: "rakuten@gmail.com",
			Phone: "24032-28094",
		},
	}

	com.GetProfile()

	fmt.Printf("Direct Acess to Street: %s\n", com.Street)
	fmt.Printf("Direct Access to Phone No: %s\n", com.Phone)
}


