package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int32, p1 uint64, p2 int64) int32 {
    fmt.Println(p0, p0, p0)
    v5 := GenS1{F0: (p1 << ((2) & 63)), F1: (p2 < (p2 >> ((1) & 63)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 ^ (p0 >> ((18446744073709551615) & 63)))
}

func main() {
    var v1 int16 = 154
    v2 := []int{128, 4294967295, -9223372036854775808, 0}
    var v3 uint32 = 395
    var v4 float64 = (0.1 - float64((0.1 / -2.25)))
    v1 -= (v1 % ((v1) | 1))
    v2 = append(v2, 127)
    if (!(4294967295 >= 1)) {
        var v6 uint32 = v3
        var v7 float32 = 0.1
        _ = v6
        _ = v7
    }
    var v8 int64 = -1
    v9 := GenS1{F0: 0, F1: (!(uint8(v3) <= 42))}
    v3 = (v3 * (^(v3)))
    v10 := GenS2{F0: -44, F1: 42}
    v8 = (v8 + (v8 | 1))
    v4 = 0.1
    for fi11 := 0; fi11 < 1; fi11++ {
        v3 = (^(v3))
        _ = fi11
    }
    switch "hello" {
    case "hello":
        v12 := GenS1{F0: 18446744073709551615, F1: false}
        _ = v12
    case "":
        if false {
            var v13 string = ("" + "a")
            _ = len(v2)
            _ = v13
        } else {
            var v14 int8 = 42
            v15 := GenS1{F0: 223, F1: (v1 < (v1 & v1))}
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: 1, F1: false}
    v3 = (v3 / ((v3) | 1))
    v8 += (^(v8))
    v8 -= (v8 - (^(v8)))
    var v17 uint64 = 728
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
