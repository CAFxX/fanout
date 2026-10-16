package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 int64, p1 uint64) int64 {
    v5 := []int{}
    var v6 int32 = int32(p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 = (p0 << ((42) & 63))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 << ((2) & 63))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := []uint32{325, 55, 2147483647}
    v4 := make(map[string]uint32)
    var v8 uint = 18446744073709551615
    var v9 int64 = 9223372036854775807
    v3 = append(v3, 2147483648)
    v3 = append(v3, 4294967295)
    delete(v1, ("foo bar" + "foo bar"))
    if false {
        if false {
            var v10 int = int(v9)
            var v11 uint = (v8 | (v8 * v8))
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(fi12)
                fmt.Println(v9, v8)
                fmt.Println(v9, v8)
                _ = fi12
            }
            v8 = 0
        }
        var v13 float32 = float32(1e10)
        if true {
            delete(v2, 506)
            v14 := GenS1{F0: 1, F1: (v8 >> ((v8) & 63))}
            _ = v14
        }
        _ = v13
    }
    _ = len(v3)
    switch 42 {
    case 128:
        v15 := GenS1{F0: 1, F1: v8}
        _ = v15
    case -729:
        v16 := GenS1{F0: 919, F1: (v8 * v8)}
        _ = v16
    case -496:
        var v17 float64 = float64(((0.1 / -2.25) + 1.5))
        _ = v17
    default:
        fmt.Println(v9, v9)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
