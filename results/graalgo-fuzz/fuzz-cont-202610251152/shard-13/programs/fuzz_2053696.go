package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }

func main() {
    var v1 int32 = -875
    var v2 int16 = 1000
    var v3 int32 = (-(v1))
    if true {
        fmt.Println(v2, v1)
        v1 -= v3
    } else {
        var v4 int8 = 104
        v4 += v4
        var v5 uint = 810
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 4294967295, F1: 0, F2: true}
    v7 := GenS1{F0: 42, F1: 18446744073709551615, F2: false}
    if false {
        fmt.Println(v3)
    } else {
        fmt.Println(v3, v3, v3)
    }
    fmt.Println(v2, v1)
    var v8 uint = 18446744073709551615
    v2 = (v2 | v2)
    var v9 bool = ((v8 % ((v8) | 1)) < (v8 >> (((v8 + (^(v8)))) & 63)))
    var v10 bool = (654 >= 255)
    var v11 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
