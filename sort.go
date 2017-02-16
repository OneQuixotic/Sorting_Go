package main	

import "fmt"
import "sort"

func main() {
    strs := []string{"d", "v", "k", "e", "j"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
	
    ints := []int{7, 2, 4, 6, 9, 11}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}