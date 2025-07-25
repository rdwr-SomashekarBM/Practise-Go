package main

import (
	"errors"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

//type Person struct {
//	name    string
//	age     int
//	ssn     int
//	ph_no   int
//	address string
//}
//
//func (person Person) getDetails() Person {
//	return person
//}
//
//func getName(person Person) string {
//	return person.name
//}
//
//func getAge(person Person) int {
//	return person.age
//}
//
//func getSsn(person Person) int {
//	return person.ssn
//}
//
//func getAddress(person Person) string {
//	return person.address
//}
//
//func getPhNo(person Person) int {
//	return person.ph_no
//}
//
//func setName(person *Person, name string) {
//	person.name = name
//}
//
//func setAge(person *Person, age int) {
//	person.age = age
//}
//
//func setSsn(person *Person, ssn int) {
//	person.ssn = ssn
//}
//
//func setAddress(person *Person, address string) {
//	person.address = address
//}
//
//func setPhNo(person *Person, ph_no int) {
//	person.ph_no = ph_no
//}

// Gin Practise:

// middleware for error handling in Gin

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err.Error(),
			})
		}

		// Any other steps if no errors are found
	}
}

func CreateServer() {
	router := gin.Default()
	router.Use(ErrorHandler())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/error", func(c *gin.Context) {
		somethingWentWrong := true

		if somethingWentWrong {
			c.Error(errors.New("something went wrong"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	router.Run(":8080")

}

func main() {
	CreateServer()
}
