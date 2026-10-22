package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }

func main() {
    var v1 int = 127
    var v2 bool = false
    var v3 int8 = -103
    v1 = (v1 & -954)
    v3 = v3
    fmt.Println(v2, v3, v3)
    var v4 float64 = -2.25
    v5 := GenS1{F0: 790, F1: 59, F2: 1}
    fmt.Println(v1, v2, v4)
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 756
        _ = fi8
        _ = v9
    }
    fmt.Println(v3)
    v10 := GenS1{F0: 290, F1: 65535, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
