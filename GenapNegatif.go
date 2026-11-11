package main
import "fmt"
func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan % 2 == 0 {
		fmt.Println("genap negatif") 
	}

	if bilangan >= 0 || bilangan % 2 != 0 {
			fmt.Println("bukan") 
		}
	}
