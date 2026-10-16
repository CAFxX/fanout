package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int }

func main() {
    v1 := []uint32{1, 506, 2, 709}
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint = 65535
        for fi6 := 0; fi6 < 10; fi6++ {
            v3 = 46
            v7 := GenS1{F0: fi4}
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v1 = append(v1, 2147483647)
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: -2147483648}
        _ = fi9
        _ = v10
    }
    if ("a" > ((("\n\t" + "hello") + "\n\t") + "\n\t")) {
        if (false && false) {
            v3 = (v3 % ((2) | 1))
            v3 = v3
        } else {
            _ = len(v2)
            v11 := GenS1{F0: 128}
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v11
        }
        var v12 float64 = (1e10 + 1e10)
        var v13 int8 = 26
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: -888}
        var v15 float32 = (0.1 + float32(0.1))
        _ = v14
        _ = v15
    }
    switch 0 {
    case 32767:
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint8 = 128
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v18 uint64 = 0
            _ = fi16
            _ = v17
            _ = v18
        }
    case 4294967295:
        v19 := GenS2{F0: 776, F1: -863, F2: 32767}
        _ = v19
    }
    var v20 int8 = 40
    var v21 uint8 = 127
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            fmt.Println(v21, fi23, v21)
            v3 = v3
            v2[fi23] = false
            _ = fi23
        }
        _ = fi22
    }
    var v24 float64 = ((0.1 + 1e10) / 1.5)
    v25 := GenS3{F0: v20, F1: (v20 % ((v20) | 1)), F2: -2147483648}
    v26 := GenS1{F0: 1}
    var v27 int = 2147483647
    switch 2147483647 {
    case -486:
        switch "a" {
        case "foo bar":
            v24 += ((v24 / 1.0) / (v24 * v24))
        }
    case -210:
        v28 := GenS2{F0: 354, F1: -32768, F2: 0}
        _ = v28
    default:
        v3 -= (^(v3))
    }
    v3 -= v3
    var v29 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v8
}
