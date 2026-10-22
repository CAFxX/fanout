package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0(p0 int32) uint {
    p0 += 0
    p0 += p0
    p0 = (p0 / (((^(p0))) | 1))
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 int16 = 580
    var v4 int8 = 0
    var v5 float32 = 0.0
    v6 := GenS1{F0: 861, F1: v4, F2: true}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int16 = (v3 | v3)
            v9 = 42
            v10 := []int{(fi8 + fi8), (fi8 & -715), (-(fi7)), (fi7 / (((fi7 << ((4294967295) & 63))) | 1))}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: 0, F1: -92, F2: true}
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: -128, F1: int8(fi12)}
        _ = fi12
        _ = v13
    }
    if (18446744073709551615 < 42) {
        delete(v2, "hello")
        v14 := []int{}
        _ = v14
    } else {
        _ = len(v1)
        v5 = 0.1
    }
    v15 := []int{-656, 127}
    for fi16 := 0; fi16 < 17; fi16++ {
        if true {
            var v17 uint8 = 0
            v18 := GenS3{F0: 0}
            _ = v17
            _ = v18
        } else {
            var v19 uint16 = 0
            var v20 string = "\n\t"
            fmt.Println(v5)
            _ = v19
            _ = v20
        }
        switch fi16 {
        case -2147483648:
            var v21 uint16 = 355
            _ = v21
        }
        if true {
            v22 := GenS4{F0: 1, F1: v3, F2: 128}
            _ = v22
        } else {
            v1[4294967295] = false
        }
        _ = fi16
    }
    if true {
        for fi23 := 0; fi23 < 11; fi23++ {
            v5 += v5
            var v24 int32 = int32(v4)
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 int8 = (v4 * 98)
            var v27 uint16 = 65535
            v28 := GenS1{F0: (v27 / ((221) | 1)), F1: -1, F2: (fi25 < fi25)}
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        var v29 float64 = (((0.0 - (1.5 + (1.5 / (0.0 + float64(((v5 - 0.0) - (v5 + v5))))))) + 1e10) * (-2.25 + 0.0))
        _ = v29
    } else {
        switch "foo bar" {
        case "":
            v30 := GenS5{F0: 40, F1: genFn0(0)}
            _ = v30
        case "foo bar":
            if true {
                v31 := GenS1{F0: 1, F1: (v4 - v4), F2: (v5 >= v5)}
                _ = v31
            }
        default:
            var v32 uint16 = 2
            _ = v32
        }
    }
    v3 = v3
    v33 := GenS6{F0: v3, F1: 2147483647, F2: 2}
    for fi34 := 0; fi34 < 17; fi34++ {
        fmt.Println(v4, v4, v5)
        _ = fi34
    }
    switch (("a" + "") + "\n\t") {
    case "hello":
        v3 = v3
    case "a":
        v4 -= (-(v4))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
