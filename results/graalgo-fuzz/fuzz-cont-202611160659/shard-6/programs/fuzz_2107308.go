package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int, p1 int8) int {
    v5 := []int{(p0 ^ p0)}
    switch p0 {
    case 2147483647:
        var v6 bool = (3.14159 >= float64((1e10 - (0.1 + 3.14159))))
        _ = v6
    case 123:
        switch p0 {
        case 0:
            if len(v5) > 0 {
                v5[0] = p0
            }
        case 796:
            v7 := GenS1{F0: 55, F1: (p0 % ((-692) | 1))}
            _ = v7
        default:
            fmt.Println(p1)
        }
    case -1:
        if true {
            var v8 float32 = 3.14159
            _ = v8
        } else {
            if len(v5) > 0 {
                v5[0] = p0
            }
            v5 = append(v5, 32768)
        }
    default:
        var v9 uint32 = 2
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: p0}
    v11 := GenS1{F0: 255, F1: (p0 / (((p0 * (-(p0)))) | 1))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v5
    return p0
}

func main() {
    v1 := []uint32{2147483647}
    var v2 float32 = float32(-2.25)
    var v3 uint16 = 1
    var v4 int16 = 534
    if false {
        v4 += (v4 << ((42) & 63))
        for fi12 := 0; fi12 < 10; fi12++ {
            v3 -= 83
            _ = fi12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS2{F0: 18446744073709551615}
            fmt.Println(v2, v2, v4)
            var v15 int64 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        if (!(994 != 9223372036854775807)) {
            var v16 int8 = -1
            _ = v16
        }
    }
    if (true || (!(false && true))) {
        var v17 uint16 = 32767
        var v18 float64 = -2.25
        _ = v17
        _ = v18
    }
    v2 -= float32(3.14159)
    _ = len(v1)
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(v4, v3)
        if len(v1) > 0 {
            v1[0] = 821
        }
        var v20 uint = 18446744073709551615
        _ = fi19
        _ = v20
    }
    v4 += (v4 << ((438) & 63))
    var v21 bool = false
    var v22 uint64 = 18446744073709551615
    var v23 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
