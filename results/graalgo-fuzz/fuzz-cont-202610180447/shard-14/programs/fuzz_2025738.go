package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 32767
    var v2 uint8 = 0
    v3 := []uint32{134}
    var v4 int8 = 0
    var v5 float32 = (1.5 / 1.0)
    for fi6 := 0; fi6 < 14; fi6++ {
        v3 = append(v3, 65535)
        _ = fi6
    }
    v7 := GenS1{F0: 65535, F1: 66}
    v2 += v2
    var v8 uint = 0
    if ((v8 ^ (v8 * (v8 ^ 255))) == (v8 >> (((v8 << ((1) & 63))) & 63))) {
        var v9 float32 = (1e10 - v5)
        var v10 uint16 = uint16(v8)
        v11 := GenS2{F0: uint8(v10), F1: ""}
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 int32 = 0
        v8 = (v8 | 0)
        switch "" {
        case "a":
            v13 := GenS1{F0: 4294967295, F1: (-(v12))}
            _ = v13
        case "hello":
            fmt.Println(v8, v8)
        case "":
            var v14 float32 = 1.5
            _ = v14
        default:
            v15 := GenS1{F0: 65535, F1: (v12 + (v12 & v12))}
            _ = v15
        }
        _ = v12
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v5 = v5
        v17 := GenS3{F0: v4}
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v2, v1)
    var v18 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
