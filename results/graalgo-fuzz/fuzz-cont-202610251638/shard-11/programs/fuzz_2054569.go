package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    var v1 int8 = -128
    var v2 bool = true
    var v3 int8 = 1
    v4 := GenS1{F0: 0, F1: 1000}
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int16 = -1
        v6 = (v6 / ((v6) | 1))
        v6 = v6
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 167, F1: 32767}
    var v8 int8 = v3
    v8 = (^(v3))
    v8 = (v1 >> ((42) & 63))
    v9 := GenS1{F0: 155, F1: 0}
    v10 := GenS2{F0: -32768, F1: 42}
    if v2 {
        var v11 bool = v2
        _ = v11
    } else {
        v1 += v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
