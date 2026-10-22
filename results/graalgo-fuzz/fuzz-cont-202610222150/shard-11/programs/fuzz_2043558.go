package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    v3 := make(map[int]bool)
    v4 := []int{-1}
    v2 -= v2
    var v5 string = ("a" + "")
    v2 = (^(v2))
    delete(v3, -128)
    v3[-9223372036854775808] = (!(("" <= ((v5 + (v5 + v5)) + v5)) && false))
    v6 := GenS1{F0: 1}
    _ = len(v4)
    v2 -= (v2 >> ((4294967295) & 63))
    v5 = v5
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
