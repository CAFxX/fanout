package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
func genFn0() int8 {
    var v4 uint8 = 28
    var v5 float32 = (-2.25 - (1e10 - 1e10))
    switch "foo bar" {
    case "":
        switch "" {
        case "a":
            v6 := []int{826}
            _ = v6
        default:
            var v7 bool = (v5 < v5)
            _ = v7
        }
    case "hello":
        var v8 int16 = -32768
        _ = v8
    default:
        fmt.Println(v4, v4)
    }
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 2
    var v3 uint32 = 2147483647
    var v9 int8 = -66
    var v10 uint8 = 2
    switch 0 {
    case 343:
        v11 := GenS1{F0: -2147483648, F1: uint(v2)}
        _ = v11
    default:
        v12 := []int{}
        _ = v12
    }
    v9 -= 42
    var v13 bool = (v1 > (v1 / (((v1 + (v1 + (v1 | v1)))) | 1)))
    v9 = v9
    switch -2147483648 {
    case -598:
        v13 = (2 > 1)
    }
    v3 = (v3 * v3)
    v14 := GenS2{F0: -478, F1: 0, F2: -9223372036854775808}
    var v15 uint32 = 1
    var v16 uint16 = (v2 * 964)
    var v17 float32 = ((float32((float64(0.0) - 1e10)) / float32((0.0 + 1e10))) + (0.0 / float32((-2.25 + 3.14159))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
