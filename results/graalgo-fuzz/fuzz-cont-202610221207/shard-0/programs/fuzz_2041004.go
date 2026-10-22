package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint8 = 128
    v2 := []int{42}
    var v3 uint16 = 42
    var v4 bool = false
    var v5 int16 = 42
    v2 = append(v2, 2147483647)
    v6 := GenS1{F0: (^(v5))}
    v7 := GenS2{F0: v4, F1: int8(v5)}
    v8 := GenS1{F0: (v5 >> ((255) & 63))}
    var v9 uint = 2
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v5, v5)
        _ = fi10
    }
    v9 = 598
    if false {
        var v11 int64 = -1
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: (v5 + (v5 % ((v5) | 1)))}
            _ = fi12
            _ = v13
        }
        var v14 int64 = 42
        _ = v14
    }
    v15 := GenS1{F0: v5}
    v16 := GenS1{F0: (v5 + -677)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
