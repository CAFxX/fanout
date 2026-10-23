package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0(p0 int32, p1 bool, p2 uint64) uint64 {
    var v6 uint32 = 65535
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 * p2)
}

func main() {
    var v1 bool = (8 > 2)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.1
    v4 := []uint32{2147483647}
    v5 := []uint32{65535, 65535, 2, 630}
    var v7 uint32 = 1
    v1 = (!true)
    fmt.Println(v7)
    v8 := v2[(^(v7))]
    if len(v5) > 0 {
        v5[0] = v8
    }
    v3 -= v3
    v9 := GenS1{F0: 127, F1: 2, F2: 42}
    var v10 int8 = 42
    if (0.0 >= (1e10 - (float32(v3) * (-2.25 + (-2.25 / (3.14159 * ((((3.14159 / 1.5) + -2.25) / (1e10 / 1.5)) + 1.5))))))) {
        var v11 uint16 = 32767
        if v1 {
            fmt.Println(v8)
            v12 := GenS2{F0: 18446744073709551615, F1: 477}
            v13 := GenS2{F0: 235, F1: -9223372036854775808}
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 42, F1: 322}
            _ = v14
        }
        _ = v11
    } else {
        switch -715 {
        case -406:
            for fi15 := 0; fi15 < 14; fi15++ {
                v4 = append(v4, v7)
                v16 := GenS3{F0: 32767}
                _ = fi15
                _ = v16
            }
        case -1:
            for fi17 := 0; fi17 < 3; fi17++ {
                if len(v4) > 0 {
                    v4[0] = v7
                }
                _ = fi17
            }
        }
        var v18 uint64 = 18446744073709551615
        fmt.Println(v10, v7)
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS4{F0: false, F1: (-(v10))}
        _ = fi19
        _ = v20
    }
    var v21 uint32 = v7
    var v22 uint = 1
    if len(v5) > 0 {
        v5[0] = (v21 << (((v22 * (v22 | v22))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
