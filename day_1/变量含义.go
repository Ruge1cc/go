package main
import "fmt"

func hello() {
	n1 := "jjjie"
	fmt.Println(n1)
}	

var (
	s1      string = "1"
	s2      string = "2"	
)

func main() {
	var name1 string = "jiejie"
	fmt.Println(name1)

	var name2 = "jie"
	fmt.Println(name2)

	name3 := "jj"
	fmt.Println(name3)

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	hello()
	fmt.Println(s1, s2)
	
}
