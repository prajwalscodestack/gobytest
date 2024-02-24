package helloworld
import(
 "fmt"
)
func SayHello() string {
	return "Hello World"
}
func SayHelloToPerson(person string) string {
	if person == "" {
		return "Hello World"
	}
	return "Hello " + person
}
func GreetEveryOne(){
 fmt.Println("Hello Everyone!")
}
