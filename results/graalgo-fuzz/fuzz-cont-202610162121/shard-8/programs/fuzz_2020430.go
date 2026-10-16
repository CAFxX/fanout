package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 uint, p1 uint64) uint {
    var v7 float32 = 0.0
    v8 := GenS1{F0: 100, F1: 9223372036854775807}
    for fi9 := 0; fi9 < 7; fi9++ {
        p1 = (^(p1))
        _ = fi9
    }
    p0 = p0
    switch 9223372036854775807 {
    case -112:
        v7 -= (((((3.14159 * 0.0) + v7) * (v7 * v7)) * float32((((-2.25 - (-2.25 - 3.14159)) * float64(1.5)) + -2.25))) - v7)
    case 828:
        var v10 float64 = (-2.25 / 3.14159)
        _ = v10
    default:
        var v11 bool = true
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, 121, 0}
    var v3 float32 = float32(1e10)
    var v4 uint64 = 742
    v5 := []uint32{633}
    v6 := []uint32{2147483647}
    v4 = 0
    v4 = 653
    switch -458 {
    case 9223372036854775807:
        for fi12 := 0; fi12 < 9; fi12++ {
            v3 = v3
            _ = fi12
        }
    case 1:
        var v13 bool = (-2147483648 != 32768)
        _ = v13
    case 42:
        fmt.Println(v3, v4)
    default:
        var v14 uint32 = 42
        _ = v14
    }
    _ = len(v1)
    if (!false) {
        fmt.Println(v3)
        var v15 string = ""
        _ = v15
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "" {
        case "hello":
            var v16 int16 = 834
            _ = v16
        case "":
            v17 := GenS1{F0: 100, F1: -1}
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int16 = 0
            _ = fi18
            _ = v19
        }
    }
    var v20 int64 = 1
    v21 := GenS1{F0: -128, F1: (-(v20))}
    v4 = (^(v4))
    if false {
        var v22 uint16 = 65535
        _ = v22
    }
    var v23 uint32 = 262
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
