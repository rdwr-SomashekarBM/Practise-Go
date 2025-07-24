package main

//type Person struct {
//	getName func() string
//	getAge  func() string
//}
//
//type PersonBuilder struct {
//	name string
//	age  string
//}
//
//func (p PersonBuilder) getName() string {
//	return p.name
//}
//
//func (p PersonBuilder) getAge() string {
//	return p.age
//}
//
//func PrintDetails(p Person) {
//	fmt.Printf("Name: %s, Age: %s\n", p.getName(), p.getAge())
//}
//
//func main() {
//	builder := PersonBuilder{name: "Soma", age: "22"}
//	p := Person{
//		getName: builder.getName,
//		getAge:  builder.getAge,
//	}
//	PrintDetails(p)
//}

//func toCheckDefer(name string) string {
//	defer fmt.Printf("defer: %s\n", name)
//	fmt.Printf("func: %s\n", name)
//	return name
//}

type Person struct {
	name    string
	age     int
	ssn     int
	ph_no   int
	address string
}

func (person Person) getDetails() Person {
	return person
}

func getName(person Person) string {
	return person.name
}

func getAge(person Person) int {
	return person.age
}

func getSsn(person Person) int {
	return person.ssn
}

func getAddress(person Person) string {
	return person.address
}

func getPhNo(person Person) int {
	return person.ph_no
}

func setName(person *Person, name string) {
	person.name = name
}

func setAge(person *Person, age int) {
	person.age = age
}

func setSsn(person *Person, ssn int) {
	person.ssn = ssn
}

func setAddress(person *Person, address string) {
	person.address = address
}

func setPhNo(person *Person, ph_no int) {
	person.ph_no = ph_no
}

func main() {

}
