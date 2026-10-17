package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
func genFn0(p0 uint8, p1 int8, p2 uint) int8 {
    p0 -= (p0 / ((p0) | 1))
    fmt.Println(p0, p0, p1)
    v5 := []int{-2147483648, 976, -2147483648, 0}
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 >> ((p2) & 63))
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = 42
    var v3 int8 = int8(v1)
    var v4 bool = true
    v3 -= genFn0(1, v3, 2)
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v2, v1, v1)
        if (!(!v4)) {
            v1 = -32768
        }
        _ = fi6
    }
    v7 := []int{718}
    fmt.Println(v2, v1, v4)
    if false {
        if (v4 && false) {
            if v4 {
                v1 = (v1 >> ((2) & 63))
                v4 = (!(253 <= 255))
            }
        }
        v2 = 42
    } else {
        var v8 int16 = v1
        switch 32767 {
        case -9223372036854775808:
            for fi9 := 0; fi9 < 20; fi9++ {
                v8 -= v1
                _ = fi9
            }
        case 0:
            fmt.Println(v3)
        default:
            fmt.Println(v4, v2, v4)
        }
        _ = len(v7)
        _ = v8
    }
    var v10 int32 = 656
    v1 = (v1 + (v1 ^ v1))
    if v4 {
        if (!(int(v2) == 32767)) {
            var v11 bool = (4294967295 < 1)
            _ = v11
        } else {
            v12 := GenS1{F0: 651, F1: 24, F2: 0}
            _ = v12
        }
        v1 = (-(v1))
    } else {
        var v13 float32 = (1.5 * 1e10)
        fmt.Println(v13, v3, v2)
        v7 = append(v7, -9223372036854775808)
        _ = v13
    }
    fmt.Println(v1, v4, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
