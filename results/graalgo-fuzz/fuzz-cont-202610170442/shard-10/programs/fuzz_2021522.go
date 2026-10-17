package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8) uint16 {
    var v5 uint = 18446744073709551615
    p0 = 2
    if false {
        var v6 uint8 = (p0 | p0)
        v7 := GenS1{F0: -95, F1: 1}
        fmt.Println(p0)
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1() uint8 {
    v8 := GenS2{F0: "hello", F1: (-1 > 2147483647)}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int32 = 42
        _ = fi9
        _ = v10
    }
    _ = v8
    return 217
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint8 = 230
    var v3 float32 = 1.5
    var v4 float64 = -2.25
    switch -1 {
    case -523:
        var v11 float32 = v3
        _ = v11
    case 901:
        fmt.Println(v2, v3, v2)
    case 1:
        v4 = (float64(v3) * v4)
    default:
        for fi12 := 0; fi12 < 2; fi12++ {
            v3 += -2.25
            for fi13 := 0; fi13 < 20; fi13++ {
                v3 = float32((v4 * 1.5))
                switch 32767 {
                case 128:
                    _ = len(v1)
                case 42:
                    v3 -= v3
                case 4294967295:
                    v3 = (1.5 + float32(v4))
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    var v14 uint64 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((true || (32767 == -962)) && ((!(false || true)) || true)) {
        v3 += (((v3 - (v3 * (3.14159 + 1e10))) / v3) / float32(v4))
    }
    v2 -= genFn1()
    if false {
        _ = len(v1)
    }
    v2 -= (v2 << ((65535) & 63))
    fmt.Println(v14, v4, v14)
    switch 42 {
    case 32767:
        switch "foo bar" {
        case "hello":
            v15 := GenS3{F0: 0}
            _ = v15
        case "":
            _ = len(v1)
        case "foo bar":
            fmt.Println(v3)
        }
    case 42:
        fmt.Println(v2, v2)
    case 32768:
        var v16 int8 = 0
        _ = v16
    default:
        var v17 uint32 = 4294967295
        _ = v17
    }
    var v18 int = 934
    var v19 float32 = (0.1 / v3)
    var v20 bool = (42 > 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
