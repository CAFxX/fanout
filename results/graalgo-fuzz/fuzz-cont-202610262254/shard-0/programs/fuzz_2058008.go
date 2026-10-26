package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int) int16 {
    if true {
        v7 := []int{-73}
        switch p0 {
        case 4294967295:
            p0 = p0
        case 2147483647:
            p0 = (p0 >> ((807) & 63))
        default:
            if ((!(0 == 896)) && (2 <= 1)) {
                fmt.Println(p0)
            }
        }
        _ = v7
    }
    if (false || (92 == 2)) {
        if (!(42 < 42)) {
            var v8 float32 = -2.25
            v9 := []int{2147483647}
            _ = v8
            _ = v9
        }
        v10 := []int{p0, (p0 ^ p0), (p0 * p0)}
        _ = v10
    }
    if true {
        if (true && (751 == 65535)) {
            var v11 uint16 = 42
            _ = v11
        }
        p0 -= (p0 * 0)
    } else {
        var v12 int = (-(p0))
        _ = v12
    }
    var v13 string = "\n\t"
    v14 := GenS1{F0: 127, F1: 128, F2: 42}
    _ = p0
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 uint = 42
    var v2 int = 32768
    var v3 uint64 = 229
    v4 := make(map[int]bool)
    var v5 uint64 = (v3 & v3)
    var v6 int32 = 1
    v3 -= (v5 | 18446744073709551615)
    v2 -= 2147483647
    _ = len(v4)
    var v15 uint = (v1 / (((v1 + (v1 << (((v1 * uint(v3))) & 63)))) | 1))
    v16 := GenS1{F0: -1, F1: 9223372036854775807, F2: 1}
    var v17 uint8 = 0
    v2 -= v2
    var v18 int = 32
    v3 = (v5 * v3)
    v4[127] = (v3 <= (v5 >> (((v1 ^ v1)) & 63)))
    var v19 string = ("foo bar" + "foo bar")
    if true {
        if false {
            v20 := []int{v18}
            v21 := GenS2{F0: (true || true)}
            _ = v20
            _ = v21
        } else {
            v6 = (v6 & (v6 ^ (v6 ^ 1)))
        }
        v1 = v1
    } else {
        v22 := GenS2{F0: (4294967295 >= uint32(v15))}
        var v23 uint64 = v5
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
