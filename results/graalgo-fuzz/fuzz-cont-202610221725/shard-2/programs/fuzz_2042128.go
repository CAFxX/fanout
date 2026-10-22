package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0() uint64 {
    v6 := []int{32767}
    switch 209 {
    case -156:
        var v7 int = 669
        _ = v7
    case 1:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case 127:
        var v9 uint8 = 127
        _ = v9
    default:
        var v10 float64 = 0.1
        _ = v10
    }
    var v11 int32 = 2147483647
    v6 = append(v6, 32768)
    v11 = v11
    _ = v11
    _ = v6
    return 42
}
func genFn1(p0 uint32) bool {
    v12 := GenS1{F0: 42, F1: -2147483648}
    var v13 uint = 942
    v14 := GenS2{F0: (p0 * 42)}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return (!(!true))
}

func main() {
    var v1 int16 = 498
    v2 := []uint32{2147483648, 65535}
    var v3 bool = (!false)
    v4 := make(map[string]uint32)
    var v5 float64 = 0.0
    if v3 {
        var v15 uint16 = 42
        _ = v15
    }
    v16 := GenS3{F0: v3, F1: 731}
    v1 = v1
    v3 = (43 != 128)
    fmt.Println(v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 float64 = (1e10 / (float64(0.1) / v5))
        switch "\n\t" {
        case "":
            v18 -= (float64(0.1) * float64(-2.25))
        }
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (true || v3) {
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS2{F0: 2147483647}
                v3 = false
                _ = fi20
                _ = v21
            }
        }
        _ = len(v4)
        var v22 string = "a"
        _ = fi19
        _ = v22
    }
    switch "" {
    case "\n\t":
        var v23 uint = 569
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
