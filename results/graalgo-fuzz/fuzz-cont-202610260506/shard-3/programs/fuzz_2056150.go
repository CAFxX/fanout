package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 string }
func genFn0() int8 {
    v4 := []int{4294967295, -2147483648}
    v5 := GenS1{F0: 100, F1: 1, F2: 32767}
    var v6 uint = 2
    v6 = (v6 / ((v6) | 1))
    _ = v4
    _ = v5
    _ = v6
    return -54
}
func genFn1() uint16 {
    var v7 int64 = 9223372036854775807
    v8 := []int{-2147483648}
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = -1
    var v3 int8 = 127
    var v9 int8 = v3
    var v10 int32 = -2147483648
    v11 := GenS1{F0: (-(v9)), F1: v10, F2: (v2 % (((v2 % (((v2 / (((v2 | 32767)) | 1))) | 1))) | 1))}
    v12 := GenS2{F0: "foo bar"}
    v1 -= 2
    var v13 uint = 0
    switch "hello" {
    case "hello":
        fmt.Println(v2, v9)
    case "foo bar":
        fmt.Println(v10, v1)
    }
    v2 = (v2 >> (((v13 + (v13 & v13))) & 63))
    v3 = (v3 - (v9 ^ (v9 & v3)))
    if (!false) {
        switch "" {
        case "foo bar":
            v1 = (^(v1))
        case "\n\t":
            v9 -= int8(v1)
        default:
            v9 = genFn0()
        }
        var v14 int32 = (v10 | (v10 ^ v10))
        v2 -= (v2 * 180)
        _ = v14
    } else {
        if ("foo bar" != "foo bar") {
            v15 := GenS1{F0: v9, F1: (v10 % ((v10) | 1)), F2: v2}
            _ = v15
        } else {
            v3 = genFn0()
            v9 = (v9 / ((genFn0()) | 1))
        }
        v13 = uint(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
