package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int8 = -39
    var v2 uint64 = 865
    var v3 uint8 = 21
    if (!false) {
        v4 := []int{-2147483648, -1}
        v5 := GenS1{F0: v1}
        _ = v4
        _ = v5
    } else {
        v6 := []int{-2147483648, 9223372036854775807, 929, -2147483648}
        v7 := GenS1{F0: 127}
        _ = v6
        _ = v7
    }
    if (!true) {
        v8 := []int{205}
        fmt.Println(v1)
        _ = v8
    }
    v3 = v3
    v9 := GenS1{F0: v1}
    var v10 int8 = (v1 + v1)
    if true {
        v10 = (-(v10))
        fmt.Println(v1, v10)
    }
    var v11 int8 = (v10 ^ (v10 | v1))
    var v12 uint8 = v3
    var v13 uint8 = v3
    v13 += 127
    v3 = v12
    var v14 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
