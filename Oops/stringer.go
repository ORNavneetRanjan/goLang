package main

import "fmt"

type School interface {
	Uniform() string
	Expertise(string, int)
}

type Student struct {
	ID     int
	Name   string
	Classs int
}

type Teacher struct {
	ID      int
	Name    string
	Subject string
	Class   []int
}

func (s Student) Uniform() string {
	return "Blue shirt and grey pants"
}

func (t Teacher) Uniform() string {
	return "Formal"
}

func (t *Teacher) Expertise(subject string, rating int) {
	if rating >= 4 {
		t.Subject = subject
	} else {
		fmt.Printf("Sorry we are not taking rating below 4/5 and yours is %d\n", rating)
	}
}

func assignSubject(school School, subject string, rating int) {
	school.Expertise(subject, rating)
}

func printUnifroms(school School){
	fmt.Println(school.Uniform())
}

func (t Teacher) String() string {
	return fmt.Sprintf("Teacher Id:%d\n Name: %s\n Expertize: %s\n Class: %v",
	 t.ID, t.Name, t.Subject, t.Class)
}

func main3() {
	t1 := Teacher{
		ID:   1,
		Name: "Navneet",
	}

	s1 := Student{
		ID:     1,
		Name:   "Ranjan",
		Classs: 10,
	}
	t1.Class = append(t1.Class, 6, 7, 8, 9, 10)

	// Pass pointer (&t1) so *Teacher satisfies School and t1.Subject is updated
	assignSubject(&t1, "Math", 4)

	printUnifroms(&t1)
	fmt.Printf("Student Uniform: %s\n",s1.Uniform())
	fmt.Printf("Teacher Unifrom: %s\n", t1.Uniform())
	fmt.Println(t1)

	

}