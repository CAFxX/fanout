package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    var v1 int32 = -766
    v2 := []uint32{2, 656}
    var v3 uint32 = 4294967295
    var v4 int = -9223372036854775808
    var v5 uint32 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v6 := GenS1{F0: 50, F1: 1}
        _ = v6
    } else {
        v3 = 596
    }
    v7 := GenS2{F0: (v1 / ((v1) | 1))}
    fmt.Println(v5, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 uint8 = 42
    v5 -= (v5 * (v5 * 4294967295))
    switch "\n\t" {
    case "a":
        var v9 uint16 = 215
        _ = v9
    case "":
        var v10 int32 = v1
        _ = v10
    case "hello":
        v11 := GenS3{F0: (v1 * v1), F1: (v1 - v1)}
        _ = v11
    }
    v4 = v4
    var v12 int64 = 42
    v12 -= 645
    var v13 uint64 = 18446744073709551615
    v14 := GenS1{F0: -128, F1: (v13 + (v13 | (v13 & (v13 % (((v13 / ((v13) | 1))) | 1)))))}
    v5 += 65535
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
