package filterset

import "fmt"

type FilterSet interface {
	Matches(string) bool
}

func SayHello() {
	fmt.Println("Hello Go!")
}
