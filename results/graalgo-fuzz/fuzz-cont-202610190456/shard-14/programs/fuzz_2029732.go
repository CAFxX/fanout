package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }

func main() {
    var v1 int8 = 100
    var v2 int = -9223372036854775808
    var v3 int16 = 42
    if ((^(v1)) <= (v1 << ((1) & 63))) {
        v4 := []int{-477}
        _ = v4
    } else {
        v5 := []int{v2, -9223372036854775808}
        v2 -= v2
        _ = v5
    }
    v6 := []int{}
    var v7 bool = ((false && true) || false)
    _ = len(v6)
    for fi8 := 0; fi8 < 18; fi8++ {
        if len(v6) > 0 {
            v6[0] = (fi8 % (((v2 + (fi8 - (fi8 ^ v2)))) | 1))
        }
        _ = fi8
    }
    switch "hello" {
    case "a":
        v7 = false
    }
    var v9 uint = 343
    fmt.Println(v2, v3)
    var v10 string = "a"
    var v11 uint16 = 42
    fmt.Println(v11)
    v12 := GenS1{F0: (-(v1)), F1: 0, F2: v10}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v13 := GenS1{F0: v1, F1: 0, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
