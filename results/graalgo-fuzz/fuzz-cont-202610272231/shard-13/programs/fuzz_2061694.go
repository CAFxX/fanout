package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint16 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int = v1
    var v3 uint32 = 104
    fmt.Println(v2, v3, v1)
    var v4 uint = 255
    fmt.Println(v3, v4)
    if false {
        v1 = v2
        v2 = (v1 << ((518) & 63))
        var v5 float64 = 0.1
        _ = v5
    }
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 9; fi7++ {
        if (-602 > 0) {
            v8 := GenS1{F0: -1, F1: 77, F2: 255}
            _ = v8
        }
        _ = fi7
    }
    v9 := []int{(v2 ^ v1)}
    switch "\n\t" {
    case "hello":
        var v10 float32 = float32(1.5)
        _ = v10
    }
    if false {
        var v11 float32 = (-2.25 + (0.0 + 1e10))
        _ = v11
    } else {
        var v12 uint32 = 4294967295
        _ = v12
    }
    var v13 bool = (-2147483648 < -1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
