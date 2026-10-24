package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 float64 = 0.0
    var v3 uint64 = 1
    fmt.Println(v2, v2, v1)
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int16 = 42
        var v6 uint16 = 65535
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint64 = 42
    v2 = v2
    fmt.Println(v2)
    fmt.Println(v1, v3, v7)
    fmt.Println(v7, v7)
    fmt.Println(v1)
    switch -2147483648 {
    case 0:
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 = v2
            _ = fi8
        }
    default:
        v9 := GenS1{F0: -74, F1: 482}
        _ = v9
    }
    v2 = v2
    var v10 int16 = 214
    v11 := []int{42, 0}
    fmt.Println(v2, v3, v1)
    v10 -= v10
    var v12 int32 = -131
    switch "a" {
    case "a":
        v10 -= (v10 ^ v10)
    case "hello":
        v2 -= 1.5
    case "foo bar":
        v10 += (v10 + (v10 | v10))
    }
    v2 = v2
    if true {
        v10 += -32768
    } else {
        v10 += v10
        fmt.Println(v1)
        v13 := GenS1{F0: 100, F1: 377}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
