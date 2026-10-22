package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8, p1 int64, p2 uint8) int {
    var v7 uint8 = p0
    v8 := []int{-885}
    v7 += (p2 & (v7 ^ (^(v7))))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 uint8 = 156
    v3 := make(map[int]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 float64 = ((-2.25 / ((float64(3.14159) + ((float64(0.1) / -2.25) - 3.14159)) / 1.0)) + -2.25)
    v6 := []int{127, 42, 4294967295}
    fmt.Println(v5, v4)
    delete(v3, 9223372036854775807)
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 0, F1: uint8(v4), F2: 128}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: -1, F1: 0, F2: uint8(v4)}
    var v12 int8 = (v1 * (v1 & v1))
    v12 = (v12 << ((0) & 63))
    if false {
        if (true && (!(846 >= 65535))) {
            var v13 int8 = -128
            _ = v13
        } else {
            fmt.Println(v2, v2, v1)
            v6 = append(v6, genFn0(58, 0, (^(v2))))
            v14 := GenS2{F0: 312, F1: 741}
            _ = v14
        }
        switch "foo bar" {
        case "":
            var v15 int32 = 492
            _ = v15
        case "a":
            v4 = 863
        case "\n\t":
            var v16 int = 367
            _ = v16
        default:
            var v17 float64 = (v5 * v5)
            _ = v17
        }
        fmt.Println(v2)
    } else {
        v18 := v3[-304]
        _ = v18
    }
    v19 := GenS3{F0: 1}
    var v20 uint16 = 1
    v21 := GenS4{F0: v12}
    for fi22 := 0; fi22 < 7; fi22++ {
        if (0.0 < (3.14159 / (-2.25 - 1e10))) {
            _ = len(v3)
            v6 = append(v6, (-(fi22)))
            var v23 uint8 = (^(v2))
            _ = v23
        }
        var v24 int8 = v12
        _ = fi22
        _ = v24
    }
    fmt.Println(v1, v12)
    if true {
        fmt.Println(v2)
        v5 = ((1.5 - v5) + (v5 + v5))
    } else {
        var v25 bool = false
        v26 := GenS1{F0: 0, F1: 194, F2: 127}
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
