// arraysum
package main

import (
	"fmt"
	"time"
)

func main() {

	var m int
	fmt.Scanf("%d", &m)
	array := make([]int, m)
	array2 := make([]int, m)
	//var array = new([elems]int)
	for i := 0; i < 2; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				fmt.Scanf("%d", &array[j])
			}
			if i == 1 {
				fmt.Scanf("%d", &array2[j])
				array2[j] = array[j] + array2[j]
			}
		}

	}
	fmt.Println(array2)
	end := time.Since(start)
	fmt.Println("time elapsed", end)
}
