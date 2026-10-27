package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 string = "a"
    var v2 uint8 = 76
    var v3 int8 = -1
    fmt.Println(v1, v2, v2)
    v3 += v3
    var v4 uint16 = 1
    var v5 int = 42
    v1 = v1
    var v6 string = v1
    v7 := []int{}
    for fi8 := 0; fi8 < 20; fi8++ {
        v4 = 42
        v9 := GenS1{F0: v5}
        _ = fi8
        _ = v9
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
