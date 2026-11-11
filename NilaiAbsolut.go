package main
import "fmt"
func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	if bilangan < 0{
		bilangan = -bilangan
	}
	fmt.Println("nilai absolut atau mutlak dari bilangan tersebut adalah", bilangan)
}

// program NilaiAbsolut
//kamus
//   bilangan : integer
//algoritma
//   input(bilangan)
//   if bilangan < 0 then
//   bilangan = -bilangan
//   endif
//   output("nilai absolut atau mutlak dari bilangan tersebut adalah", bilangan)
//endprogram