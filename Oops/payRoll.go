package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)",
se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name					string
	AnnualSalary	float64
	HourWorked		int
	Rate					float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return float64(he.HourWorked) * he.Rate
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %d)", 
		he.Name, he.Rate, he.HourWorked)
}

type CommissionEmployee struct {
	Name						string
	BaseSalary			float64
	CommissionRate	float64
	SalesAmount			float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount) / 100
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commissioned: %s Base Salary: %.2f, Commission: %.2f, Total: %.2f",
	 ce.Name, ce.BaseSalary, ce.CommissionRate * ce.SalesAmount / 100, ce.CalculatePay())
}

func PrintEmployeeSummary[P fmt.Stringer](emp P){
	fmt.Printf(" - Processing: %s\n", emp);
}

// func PrintEmployeeSummary (emp Payable){
// 	fmt.Printf(" - Processing: %s\n", emp);
// }

func ProcessPayroll(employee []Payable){
	fmt.Println("--- Processing Payments ---")
	totalAmount :=  0.0
	for _, emp := range employee {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("__ Monthly Pay: $%.2f\n", pay)
		totalAmount += pay
	}

	fmt.Printf("\nTotal Monthly Payroll: $%.2f", totalAmount)
	fmt.Println("-----------------------------------------------")
}

func main(){

	fmt.Println("Welcome to the Payroll Processor!")

	salEmp := SalariedEmployee{Name: "Takshay", AnnualSalary: 60000}
	hrEmp := HourlyEmployee{Name: "Khyati", AnnualSalary: 35000, Rate: 40, HourWorked: 40}
	comEmp := CommissionEmployee{Name: "Adarsh", BaseSalary: 30000, SalesAmount: 100000, CommissionRate: 10}

	payrollList := []Payable {
		salEmp, hrEmp, comEmp,
		HourlyEmployee{Name: "Navneet Ranjan", HourWorked: 40, Rate: 30},
	}
	ProcessPayroll(payrollList)

}