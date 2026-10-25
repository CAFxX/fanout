package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
func genFn0(p0 uint64) int16 {
    v4 := []int{-9223372036854775808, 2147483647}
    v4 = append(v4, 0)
    var v5 int32 = -762
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 2147483647, F1: true, F2: 440}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 370
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 4294967295
    var v3 uint32 = v1
    var v8 int16 = -1
    var v9 uint32 = (v1 % ((v3) | 1))
    v8 = (v8 & (^(v8)))
    var v10 float32 = (1e10 / 1.5)
    var v11 float64 = (-2.25 * 3.14159)
    v2 += (v2 - (v2 & v2))
    fmt.Println(v11)
    switch 127 {
    case 127:
        v8 = v8
    case 1:
        v12 := []int{424}
        _ = v12
    case 42:
        fmt.Println(v10)
    default:
        var v13 int8 = -128
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 bool = ((!(((v11 * v11) * 0.0) == 0.1)) && true)
            _ = fi14
            _ = v15
        }
    }
    var v16 bool = ((!((3.14159 + 0.0) == v10)) && false)
    fmt.Println(v1)
    if false {
        if v16 {
            v17 := GenS2{F0: -825, F1: 18446744073709551615, F2: 1}
            var v18 uint8 = 42
            _ = v17
            _ = v18
        }
    } else {
        fmt.Println(v10)
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int8 = 127
            _ = fi19
            _ = v20
        }
        v8 -= (v8 * (v8 ^ v8))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
