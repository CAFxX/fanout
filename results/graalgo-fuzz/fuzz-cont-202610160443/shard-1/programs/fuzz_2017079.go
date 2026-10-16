package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    var v3 int = 90
    var v4 int64 = -1
    var v5 float32 = (float32(float64(float32(((float64(1e10) / -2.25) - 1e10)))) * float32(1.5))
    v5 += 0.0
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "a" {
        case "foo bar":
            var v7 uint32 = 2147483647
            _ = v7
        case "\n\t":
            v1 = -128
        default:
            var v8 int32 = 845
            _ = v8
        }
        v9 := []int{(fi6 | fi6), 32767}
        _ = fi6
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        var v10 float32 = 1e10
        _ = v10
    case "\n\t":
        v1 -= (v1 / ((v1) | 1))
    case "":
        var v11 int8 = (v1 + -128)
        _ = v11
    }
    v12 := GenS1{F0: "hello", F1: 1, F2: 105}
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: "a", F1: 42, F2: 127}
        v15 := GenS2{F0: v1, F1: 0, F2: (v1 >> ((2) & 63))}
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS1{F0: ("a" + ""), F1: 42, F2: 2}
            v1 = 100
            _ = len(v2)
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v18 := []int{408}
    var v19 string = ("\n\t" + "foo bar")
    _ = len(v18)
    v5 += (v5 - 3.14159)
    v3 -= (v3 + (v3 ^ v3))
    v3 += (v3 % ((int(v1)) | 1))
    if true {
        v20 := GenS2{F0: v1, F1: 42, F2: (-(v1))}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
