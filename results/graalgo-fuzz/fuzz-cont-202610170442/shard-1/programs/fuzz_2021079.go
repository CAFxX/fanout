package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 uint32) bool {
    var v6 int16 = 32767
    v7 := GenS1{F0: p2, F1: 0, F2: 128}
    var v8 int32 = 652
    var v9 uint8 = 1
    fmt.Println(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (p0 < p0)
}
func genFn1(p0 int, p1 int8, p2 int8) int8 {
    p0 = (p0 ^ (p0 * 127))
    if ("a" == "foo bar") {
        var v10 float64 = 1.5
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 string = "\n\t"
    var v5 int64 = -122
    for fi11 := 0; fi11 < 3; fi11++ {
        switch (fi11 | -106) {
        case -435:
            var v12 uint32 = 0
            _ = v12
        default:
            v4 = (v4 + "\n\t")
        }
        _ = fi11
    }
    v5 += (-(v5))
    fmt.Println(v4, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 uint8 = 162
    var v14 uint = 18446744073709551615
    v2[32767] = 2
    fmt.Println(v4)
    v15 := GenS1{F0: 565, F1: 0, F2: 128}
    fmt.Println(v4)
    var v16 uint = (v14 / ((246) | 1))
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v5 -= (-(v5))
            _ = fi18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint8 = v13
            _ = fi19
            _ = v20
        }
        var v21 int = (fi17 + (fi17 << ((v14) & 63)))
        _ = fi17
        _ = v21
    }
    var v22 int = -93
    if false {
        v5 = -646
        _ = len(v1)
        v23 := GenS1{F0: 1, F1: v16, F2: 2}
        _ = v23
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
