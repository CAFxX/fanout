package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0() int16 {
    var v7 int16 = -32768
    var v8 int32 = 42
    v7 = v7
    v9 := []int{-9223372036854775808, 537}
    v7 = v7
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint64, p1 int32) int {
    v10 := []int{-1, -1, -363}
    p1 += (p1 ^ 42)
    if false {
        var v11 float64 = (0.0 - 1.5)
        var v12 uint32 = 2147483647
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "":
        var v13 int32 = (p1 * p1)
        _ = v13
    case "foo bar":
        if false {
            p0 -= p0
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi14, fi14)
                v15 := GenS1{F0: ((1 < 718) || (-118 <= -1)), F1: 1}
                fmt.Println(p1, fi14, fi14)
                _ = fi14
                _ = v15
            }
        }
    default:
        var v16 int32 = (-(p1))
        _ = v16
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = p1
    _ = v10
    return 9223372036854775807
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    var v4 uint = 255
    var v5 uint16 = 0
    var v6 uint = (v4 << (((v4 + (v4 & 554))) & 63))
    delete(v2, 650)
    var v17 bool = (-2.25 >= 1.5)
    v18 := GenS1{F0: v17, F1: -1}
    v6 += v4
    fmt.Println(v1, v4, v5)
    v19 := GenS1{F0: v17, F1: 0}
    if (!((false && (!v17)) || ((18446744073709551615 == 18446744073709551615) || (42 > -1)))) {
        v20 := GenS2{F0: (v5 | 32767), F1: 210}
        var v21 int64 = 9223372036854775807
        switch 9223372036854775807 {
        case 127:
            fmt.Println(v1, v4, v21)
        }
        _ = v20
        _ = v21
    } else {
        if v17 {
            var v22 int = 32767
            _ = v22
        } else {
            v23 := GenS2{F0: 824, F1: -935}
            _ = v23
        }
        var v24 uint64 = 18446744073709551615
        var v25 float64 = 3.14159
        _ = v24
        _ = v25
    }
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
