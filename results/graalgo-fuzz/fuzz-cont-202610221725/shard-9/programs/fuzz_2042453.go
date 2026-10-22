package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }

func main() {
    var v1 uint = 168
    var v2 float64 = (1e10 / 1e10)
    var v3 int64 = 1
    fmt.Println(v1, v3)
    fmt.Println(v1, v1, v2)
    v2 += 3.14159
    switch "a" {
    case "a":
        if false {
            var v4 float64 = (v2 * v2)
            _ = v4
        } else {
            var v5 int = 32768
            var v6 float64 = (v2 * float64(1e10))
            _ = v5
            _ = v6
        }
    case "hello":
        v3 += (v3 - v3)
    default:
        var v7 float64 = v2
        _ = v7
    }
    var v8 float32 = ((float32(v2) / 3.14159) + 0.0)
    v9 := GenS1{F0: 1000, F1: 504, F2: 2147483647}
    var v10 uint16 = 0
    v11 := GenS1{F0: -975, F1: v10, F2: 805}
    var v12 string = "hello"
    v1 += v1
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            v12 = v12
            v15 := GenS1{F0: -32768, F1: (v10 + 0), F2: uint32(v1)}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS1{F0: -32768, F1: 32767, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
