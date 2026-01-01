package main 

// import "fmt"

func Index(a,b string) int{
	count:=0
	for i:=0; i<len(a);i++{
		for j:=0;j<len(b);j++{
			if a[i] == b[j]{
				count++
			}
		}
	}
	return count
	
}
// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }