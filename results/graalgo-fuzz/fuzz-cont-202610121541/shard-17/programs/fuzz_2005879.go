package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0() bool {
    if (float64(3.14159) >= 1.5) {
        var v5 uint16 = 42
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 92, F1: -462}
        var v8 uint = 329
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int = -2147483648
    _ = v9
    return false
}
func genFn1(p0 int8, p1 uint64, p2 string) uint16 {
    v10 := GenS2{F0: 2, F1: true}
    var v11 string = (("\n\t" + ("" + ((p2 + p2) + "foo bar"))) + (p2 + "foo bar"))
    if (75 <= p0) {
        p0 -= -128
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(p0)
            for fi13 := 0; fi13 < 9; fi13++ {
                switch (fi12 | (fi12 | (fi12 / ((fi13) | 1)))) {
                case 32768:
                    if (false && false) {
                        p1 += 971
                        var v14 uint8 = 255
                        v14 += v14
                        _ = v14
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
    } else {
        v15 := []int{0, 1}
        var v16 uint16 = 65535
        _ = v15
        _ = v16
    }
    if (42 <= int16(p0)) {
        p0 -= p0
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int = (-(fi17))
            _ = fi17
            _ = v18
        }
    }
    var v19 int64 = 124
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v19
    return 884
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = -2.25
    var v3 uint16 = 952
    var v4 uint = 35
    v2 = float64(0.1)
    v20 := GenS1{F0: 74, F1: 0}
    v21 := []int{-1, -856, -2147483648}
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 string = ""
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int64 = 665
        _ = len(v21)
        fmt.Println(v4, v3)
        _ = fi24
        _ = v25
    }
    var v26 float32 = 0.0
    v1 = (^(v1))
    var v27 uint8 = 67
    v3 = 370
    _ = len(v21)
    switch "a" {
    case "a":
        switch -981 {
        case 837:
            _ = len(v21)
        default:
            var v28 int64 = 42
            _ = v28
        }
    default:
        var v29 string = "\n\t"
        _ = v29
    }
    var v30 uint8 = 2
    v4 = 2
    var v31 int32 = 343
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
