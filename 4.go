package main

import "fmt"

func main() {
 var N, i int
 var topi, alatTulis, tali, pisau, status bool

 fmt.Scanln(&N)
 status = true
 i = 0
 for i < N && status {
  fmt.Scanln(&topi, &alatTulis, &tali, &pisau)
  status = (topi && alatTulis && tali && pisau)
  i += 1

 }
 fmt.Println(status)
}