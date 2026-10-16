package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int64; F2 uint16 }
func genFn0(p0 uint8, p1 int32, p2 uint16) uint32 {
    var v6 uint16 = (p2 & p2)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{fi7, (fi7 | fi7)}
        for fi9 := 0; fi9 < 3; fi9++ {
            p2 -= p2
            var v10 uint8 = 0
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: p0}
        _ = fi7
        _ = v11
        _ = v8
    }
    p0 = 1
    var v12 uint = 328
    var v13 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 int64, p2 int8) uint32 {
    var v14 float32 = 3.14159
    var v15 uint64 = 4294967295
    v16 := []int{944}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float64 = (1.5 / 3.14159)
        fmt.Println(v14, p2, fi17)
        if (((v14 * v14) / v14) > (v14 * v14)) {
            var v19 uint32 = 2147483648
            var v20 float32 = (0.1 * 0.1)
            v16 = append(v16, -1)
            _ = v19
            _ = v20
        } else {
            v18 = (float64(0.1) * 1.5)
            v21 := GenS1{F0: 1}
            v22 := GenS2{F0: (fi17 % ((fi17) | 1))}
            _ = v21
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return genFn0(127, -799, p0)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    var v4 int16 = -1
    v5 := []uint32{65535}
    if false {
        v4 = -102
    }
    v4 += v4
    fmt.Println(v4)
    v3 -= (^(v3))
    v23 := GenS3{F0: false, F1: 42, F2: 487}
    fmt.Println(v4)
    fmt.Println(v3, v4)
    fmt.Println(v3, v4)
    fmt.Println(v4, v4)
    v5 = append(v5, genFn1(255, 42, -100))
    var v24 int64 = 1
    var v25 int16 = v4
    v4 = v25
    v26 := GenS2{F0: -942}
    if false {
        switch "foo bar" {
        case "\n\t":
            var v27 string = "foo bar"
            _ = v27
        default:
            fmt.Println(v3, v25)
        }
    }
    v5 = append(v5, 292)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
