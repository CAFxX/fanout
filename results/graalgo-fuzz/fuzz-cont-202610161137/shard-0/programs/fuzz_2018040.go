package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []uint32{4294967295, 2147483647, 47}
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int32 = 1
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        v5 -= (v5 ^ 117)
    case "foo bar":
        _ = len(v4)
    case "":
        var v8 int8 = -128
        _ = v8
    }
    var v9 uint32 = 80
    if true {
        var v10 float64 = 1.5
        v11 := GenS1{F0: v5, F1: "foo bar", F2: (((0.1 * 3.14159) - 1e10) != 0.0)}
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: -426, F1: 0}
    var v13 string = ""
    fmt.Println(v9, v5)
    v14 := GenS3{F0: -128, F1: 18446744073709551615}
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 float32 = 3.14159
        var v17 string = ("\n\t" + "")
        var v18 int8 = -128
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v5 += (v5 + 18446744073709551615)
    v19 := GenS4{F0: 657, F1: v9, F2: 113}
    v20 := v3[v13]
    switch -338 {
    case 325:
        v13 = v13
    case 357:
        v21 := GenS2{F0: 0, F1: int(v5)}
        _ = v21
    default:
        for fi22 := 0; fi22 < 19; fi22++ {
            v20 -= v20
            if true {
                switch (fi22 % ((fi22) | 1)) {
                case 0:
                    fmt.Println(v20, fi22, v20)
                case 32767:
                    var v23 int = (fi22 ^ fi22)
                    _ = v23
                case 128:
                    var v24 int16 = 1
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v3["hello"] = v9
                        _ = fi25
                    }
                }
            } else {
                if (-1 != 42) {
                    var v26 int = (fi22 | -744)
                    _ = v26
                }
            }
            v9 = 4294967295
            _ = fi22
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 = 0
    v27 := GenS2{F0: 0, F1: 4294967295}
    v28 := GenS5{F0: v9, F1: 1000}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
