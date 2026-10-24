package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 uint64 = 18446744073709551615
    var v3 string = "hello"
    v3 = (v3 + v3)
    v2 = v2
    var v4 uint8 = 73
    var v5 uint32 = 4294967295
    var v6 bool = (-2.25 > -2.25)
    switch 32767 {
    case 127:
        v7 := []int{}
        _ = v7
    }
    var v8 int16 = 1000
    if v6 {
        switch "\n\t" {
        case "":
            var v9 int32 = -2147483648
            _ = v9
        default:
            var v10 int32 = 2147483647
            _ = v10
        }
    } else {
        v11 := GenS1{F0: (v1 * 255), F1: 42, F2: -1}
        v12 := []int{-1}
        if len(v12) > 0 {
            v12[0] = -9223372036854775808
        }
        _ = v11
        _ = v12
    }
    var v13 int16 = (v8 & 32767)
    v14 := GenS2{F0: v1, F1: (v1 + 32767)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
