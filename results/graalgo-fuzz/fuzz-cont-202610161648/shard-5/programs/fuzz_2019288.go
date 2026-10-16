package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0(p0 int32, p1 int32) uint32 {
    p0 = p1
    var v7 uint64 = 4294967295
    var v8 float64 = (float64(0.0) / ((1e10 + 1e10) + -2.25))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint) uint8 {
    fmt.Println(p0)
    p0 += (p0 % ((p0) | 1))
    var v9 int8 = -128
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    v6 := make(map[int]uint32)
    var v10 int8 = 42
    v2 -= (v2 % ((v2) | 1))
    v10 += 42
    v11 := []int{520, 127}
    v5 += (-2.25 * (v5 / (v5 * v5)))
    var v12 uint32 = 444
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 float64 = float64(((0.1 + 3.14159) * 1e10))
            var v16 uint = 18446744073709551615
            v6[fi13] = (^(v12))
            v10 -= v10
            v17 := GenS1{F0: true, F1: 42}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    switch "\n\t" {
    case "":
        if (false || true) {
            v18 := GenS2{F0: true, F1: 0}
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 bool = false
                _ = fi19
                _ = v20
            }
            switch ("" + "\n\t") {
            case "foo bar":
                v21 := GenS2{F0: (!true), F1: 2147483647}
                _ = v21
            case "a":
                if (!false) {
                    v22 := GenS3{F0: 1}
                    _ = v22
                }
            case "\n\t":
                for fi23 := 0; fi23 < 11; fi23++ {
                    _ = len(v4)
                    _ = fi23
                }
            default:
                v24 := v3[""]
                _ = v24
            }
            _ = v18
        } else {
            delete(v6, 761)
        }
    default:
        v25 := GenS4{F0: 126, F1: v10}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
