package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 39
    v3 := []int{127, 548}
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int = (-(fi4))
        v1 -= v1
        fmt.Println(v5, v5, fi4)
        _ = fi4
        _ = v5
    }
    var v6 uint = 1
    var v7 string = "\n\t"
    var v8 int16 = 619
    var v9 bool = (v2 >= v2)
    v8 -= -241
    if len(v3) > 0 {
        v3[0] = 127
    }
    v10 := GenS1{F0: v2, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
