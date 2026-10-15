package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int32 = 358
    var v2 int8 = 89
    var v3 int32 = 42
    var v4 uint8 = 189
    var v5 int32 = (v1 >> ((42) & 63))
    v4 = (v4 ^ v4)
    var v6 int8 = v2
    v7 := GenS1{F0: (-1 > 9223372036854775807), F1: 2}
    switch -9223372036854775808 {
    case -364:
        if (!(3.14159 >= 0.1)) {
            var v8 int32 = v5
            v8 += int32(v4)
            _ = v8
        }
    case 2147483647:
        var v9 float64 = 0.1
        _ = v9
    default:
        v5 -= (v1 % ((2147483647) | 1))
    }
    v4 = 236
    v10 := []int{605, 32767, 1, -18}
    switch 19 {
    case 21:
        if false {
            v2 += v2
        }
    }
    v6 = (v2 / (((v2 + (v6 / ((127) | 1)))) | 1))
    var v11 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
