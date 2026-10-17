package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
func genFn0(p0 string, p1 int64) uint {
    v5 := GenS1{F0: 1, F1: (p1 ^ (^(p1))), F2: (-159 > -1)}
    v6 := []int{-1}
    v7 := GenS1{F0: 1000, F1: (p1 * (p1 * 0)), F2: (39 <= 0)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]bool)
    var v3 uint16 = 2
    v4 := []uint32{196}
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 8; fi9++ {
        v3 -= v3
        v1 = (v1 & 65535)
        v3 = 42
        _ = fi9
    }
    var v10 float32 = 0.0
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int16 = 32767
        switch int(v12) {
        case 128:
            v8 -= 2147483647
        case 0:
            var v13 string = ""
            _ = v13
        case -2147483648:
            v8 = (v8 - v8)
        default:
            v1 = (^(v1))
        }
        var v14 float32 = v10
        _ = fi11
        _ = v12
        _ = v14
    }
    fmt.Println(v10, v3)
    v8 = (v8 + -2147483648)
    var v15 float32 = v10
    var v16 float32 = (v15 - (v10 / v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
