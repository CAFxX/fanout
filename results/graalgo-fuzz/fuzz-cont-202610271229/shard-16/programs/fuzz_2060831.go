package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 string, p1 int32, p2 int8) uint64 {
    switch -9223372036854775808 {
    case 385:
        if (42 <= 429) {
            var v4 float64 = 0.1
            var v5 uint8 = 255
            v6 := []int{int(p2), 32768}
            _ = v4
            _ = v5
            _ = v6
        }
    case -589:
        p0 = (p0 + ("\n\t" + p0))
    default:
        p1 = (p1 + (p1 | -422))
    }
    var v7 uint64 = 583
    p2 = (p2 | p2)
    v8 := []int{2147483647, 2147483647}
    if (!true) {
        fmt.Println(v7, p0)
        var v9 uint32 = 42
        _ = v9
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint16 = 385
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            p1 = p1
            _ = fi10
            _ = v11
        }
        if len(v8) > 0 {
            v8[0] = 2147483647
        }
        if (!(((758 != 9223372036854775807) || (314 >= 0)) || true)) {
            switch -410 {
            case 0:
                p2 = 120
            default:
                v12 := GenS1{F0: 0, F1: 929}
                _ = v12
            }
            if (832 != 0) {
                v13 := GenS2{F0: (-(p1)), F1: -9223372036854775808, F2: 0}
                _ = v13
            } else {
                if len(v8) > 0 {
                    v8[0] = 42
                }
                v14 := GenS2{F0: -1, F1: 0, F2: 1}
                _ = v14
            }
            p1 -= p1
        } else {
            var v15 uint = 18446744073709551615
            v8 = append(v8, -534)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v7 + (v7 * (v7 ^ 4294967295)))
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]uint32)
    var v3 uint64 = 18446744073709551615
    fmt.Println(v1)
    v16 := []int{0, 0}
    v17 := GenS3{F0: -448, F1: 32767}
    v18 := v2[743]
    v3 -= (v3 + v3)
    var v19 int = -156
    v20 := GenS1{F0: 1, F1: (^(v18))}
    _ = len(v2)
    fmt.Println(v18, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
