package main

import (
    "fmt"
)

func main() {
    days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
    daySlices := days[1:3]

    daySlices = append(daySlices, "NewDay")
    fmt.Println(len(daySlices))
    fmt.Println(cap(daySlices))

    nombres:=make([]string, 5, 10)
    nombres[0] = "Juan"
    fmt.Println(nombres)

    slices1 := []int{1, 2, 3, 4, 5}
    slices2 := make([]int, 5)
    copy(slices2, slices1)
}