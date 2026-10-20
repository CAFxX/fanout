package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0() uint16 {
    var v5 uint = 18446744073709551615
    fmt.Println(v5)
    v6 := []int{1, 424}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 443, F1: 279, F2: -1}
        _ = fi7
        _ = v8
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint, p1 uint8, p2 uint32) int32 {
    var v9 int = -612
    v9 = v9
    var v10 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 float32 = ((1.5 - 0.1) / 1.0)
    var v2 uint32 = 42
    var v3 uint64 = 1
    v4 := make(map[uint32]bool)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS2{F0: v2, F1: 42}
        switch "" {
        case "":
            fmt.Println(v3, fi11, fi11)
        case "foo bar":
            var v13 string = ((("a" + "\n\t") + "foo bar") + "foo bar")
            _ = v13
        default:
            v1 -= float32(3.14159)
        }
        _ = fi11
        _ = v12
    }
    v3 -= v3
    var v14 uint = 1
    if (0 > 2) {
        v3 = 2
        for fi15 := 0; fi15 < 3; fi15++ {
            switch "" {
            case "foo bar":
                fmt.Println(fi15, fi15, v14)
            }
            _ = fi15
        }
        v14 = v14
    }
    v1 -= (v1 - 3.14159)
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := []int{42, -1}
            _ = fi17
            _ = v18
        }
        fmt.Println(v3, v2, v1)
        var v19 uint = (v14 % ((v14) | 1))
        _ = fi16
        _ = v19
    }
    var v20 int64 = 1
    v21 := GenS1{F0: 305, F1: 247, F2: -32768}
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := v4[v2]
        _ = fi22
        _ = v23
    }
    v24 := []int{1, 810, 32768}
    v1 = 0.0
    switch ("" + "hello") {
    case "\n\t":
        var v25 float64 = ((1.5 + 0.1) - 0.1)
        _ = v25
    default:
        var v26 int16 = -1
        _ = v26
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
