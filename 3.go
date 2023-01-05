package main
import "fmt"

func main(){
	var n, i, jumlah, batas int
	fmt.Scan(&n)
	i = 1
	jumlah = 0
	for i <= n{
		fmt.Scan(&batas)
		for batas < 0 || batas > 9{
			fmt.Scan(&batas)
		}
		jumlah += batas
		i += 1
	}
	fmt.Print(jumlah)
}