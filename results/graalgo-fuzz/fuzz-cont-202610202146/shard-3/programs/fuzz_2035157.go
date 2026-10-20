package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string; F1 uint32; F2 string }
func genFn0(p0 uint8) uint {
    if (1.5 <= 1e10) {
        var v7 uint16 = 1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(fi9)
                if (false || false) {
                    if false {
                        v10 := GenS1{F0: -1, F1: (fi9 % (((fi8 + (-(fi8)))) | 1)), F2: 42}
                        _ = v10
                    }
                } else {
                    fmt.Println(fi8)
                    p0 += p0
                    p0 -= p0
                }
                fmt.Println(fi8)
                _ = fi9
            }
            _ = fi8
        }
        var v11 float64 = 1.5
        var v12 int16 = -319
        _ = v11
        _ = v12
    }
    p0 = p0
    var v13 int = 2147483647
    p0 = (^(p0))
    _ = p0
    _ = v13
    return 0
}
func genFn1(p0 bool, p1 uint16) bool {
    v14 := GenS1{F0: 100, F1: 127, F2: 4294967295}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v14
    return p0
}

func main() {
    var v1 float32 = (0.0 * 0.1)
    var v2 float64 = 1e10
    v3 := make(map[int]string)
    var v4 int8 = 127
    v5 := make(map[int]bool)
    var v6 uint16 = 0
    var v15 int = 127
    v15 += v15
    v16 := []int{(v15 + v15)}
    switch -1 {
    case 128:
        for fi17 := 0; fi17 < 16; fi17++ {
            v6 = (v6 / ((1) | 1))
            v18 := GenS1{F0: (v4 + (v4 << ((255) & 63))), F1: fi17, F2: 2}
            fmt.Println(v6)
            _ = fi17
            _ = v18
        }
    default:
        v4 -= 6
    }
    if (!false) {
        v1 -= ((0.1 + 3.14159) - 0.1)
    } else {
        fmt.Println(v1)
    }
    var v19 int = (v15 & (-(v15)))
    v1 = 0.0
    v20 := GenS1{F0: int8(v6), F1: v19, F2: 4294967295}
    v21 := GenS2{F0: 32767, F1: 352}
    switch "hello" {
    case "hello":
        _ = len(v3)
    case "a":
        if true {
            var v22 bool = false
            v1 = float32(float64(0.1))
            _ = v22
        }
    case "\n\t":
        _ = len(v3)
    }
    var v23 string = "\n\t"
    v24 := GenS3{F0: "a", F1: 94, F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
