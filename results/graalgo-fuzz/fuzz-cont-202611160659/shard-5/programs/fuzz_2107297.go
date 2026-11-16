package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint, p1 int16) int16 {
    if false {
        if false {
            p1 += (p1 >> (((p0 + (^(p0)))) & 63))
        } else {
            p0 = 264
        }
    } else {
        var v5 int64 = 727
        _ = v5
    }
    p1 -= (p1 - p1)
    v6 := []int{9223372036854775807, -1}
    if len(v6) > 0 {
        v6[0] = 32767
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 uint8) int64 {
    if false {
        v7 := GenS1{F0: false, F1: 2, F2: 23}
        var v8 float32 = (((3.14159 * 0.0) / float32(float64(0.1))) + 1e10)
        _ = v7
        _ = v8
    } else {
        var v9 string = "\n\t"
        v10 := GenS1{F0: (0 > 2), F1: 624, F2: -1}
        _ = v10
        _ = v9
    }
    var v11 uint = 18446744073709551615
    v12 := []int{}
    _ = p0
    _ = v11
    _ = v12
    return 318
}

func main() {
    v1 := []uint32{42, 384, 2147483647, 2}
    var v2 float64 = 0.1
    v3 := make(map[uint32]string)
    var v4 uint64 = 4294967295
    switch 630 {
    case -614:
        var v13 uint8 = 42
        _ = v13
    case 42:
        v14 := GenS1{F0: (!(!false)), F1: 195, F2: 42}
        _ = v14
    case -334:
        v15 := GenS1{F0: false, F1: 2, F2: 9223372036854775807}
        _ = v15
    default:
        v16 := GenS2{F0: (127 != 32768), F1: 80}
        _ = v16
    }
    var v17 uint16 = 42
    var v18 int8 = 15
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint64 = (v4 / ((v4) | 1))
        _ = fi19
        _ = v20
    }
    v2 += ((((v2 * ((v2 * v2) / 1e10)) / (3.14159 * v2)) - v2) * v2)
    var v21 uint = 744
    v18 = v18
    var v22 uint16 = (v17 >> ((v21) & 63))
    v4 += v4
    var v23 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
