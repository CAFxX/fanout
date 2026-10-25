package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 float64 = float64((3.14159 * -2.25))
    var v2 int8 = 42
    v3 := make(map[uint32]string)
    var v4 uint16 = 42
    var v5 bool = (-252 < 9223372036854775807)
    var v6 int32 = 2147483647
    v6 += (v6 * 2147483647)
    if v5 {
        v7 := []int{-157, 1, 9223372036854775807, 1}
        _ = v7
    }
    fmt.Println(v4, v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint8 = 85
        if v5 {
            v4 = 70
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int = fi8
                v12 := GenS1{F0: -1, F1: -9223372036854775808}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        var v13 bool = v5
        _ = fi8
        _ = v13
        _ = v9
    }
    var v14 uint16 = 65535
    if (4294967295 > 815) {
        switch 52 {
        case 0:
            v4 -= (v4 * (v14 % (((v4 - (v4 ^ 1))) | 1)))
        case 32767:
            fmt.Println(v5, v2)
        default:
            fmt.Println(v6)
        }
    }
    v4 -= 32767
    v15 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
