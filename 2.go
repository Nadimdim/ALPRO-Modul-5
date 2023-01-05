package main

import "fmt"

func main() {
 var N, k0, k1 int
 var status bool
 fmt.Scan(&N)

 status = true

 k0 = N % 10
 N /= 10

 for N > 0 && status == true {
  k1 = N % 10
  N /= 10

  status = (k0-k1 == 1) || (k0-k1 == -1)

  k0 = k1
 }

 fmt.Println(status)
}