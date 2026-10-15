package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := []int{2147483647, 895}
    var v2 uint = 1
    var v3 int64 = 9223372036854775807
    var v4 uint = v2
    var v5 float64 = float64(((-2.25 / (1e10 + (3.14159 - (float32(float64(1e10)) + (3.14159 / (-2.25 / 0.1)))))) + (-2.25 * 1e10)))
    var v6 int8 = -12
    fmt.Println(v6, v6)
    if (!(0 == 42)) {
        fmt.Println(v6)
        v7 := GenS1{F0: (!(-1 == -32768)), F1: (v6 << (((v2 * (v4 >> (((v4 ^ (^(v2)))) & 63)))) & 63))}
        _ = v7
    }
    fmt.Println(v2)
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: true, F1: (v6 % ((v6) | 1))}
        v4 += v2
        _ = fi8
        _ = v9
    }
    var v10 float64 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 float32 = (float32(v10) - ((0.1 * 3.14159) * 1e10))
        switch "\n\t" {
        case "a":
            v13 := GenS1{F0: true, F1: (-(v6))}
            _ = v13
        case "\n\t":
            v12 = v12
        }
        _ = fi11
        _ = v12
    }
    var v14 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
