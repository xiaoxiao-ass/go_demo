package main

import "fmt"

func main() {
     sum:=0

	for i := 0; i < 10; i++ {
		if i%2==0{
			sum+=i
		}else {
			continue
		}
	}
	fmt.Println(sum)
}
