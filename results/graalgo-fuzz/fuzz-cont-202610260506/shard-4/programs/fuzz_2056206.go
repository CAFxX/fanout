package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int; F1 uint32 }

func main() {
    var v1 int64 = 42
    v2 := []uint32{}
    var v3 uint32 = 543
    if true {
        switch 2147483647 {
        case 1:
            v4 := GenS1{F0: -898, F1: -84, F2: 42}
            _ = v4
        case -834:
            var v5 uint8 = 2
            _ = v5
        default:
            v1 += v1
        }
        fmt.Println(v3)
    } else {
        v6 := GenS2{F0: 18446744073709551615, F1: 100, F2: v3}
        var v7 int8 = -128
        _ = v6
        _ = v7
    }
    if true {
        var v8 string = "foo bar"
        var v9 float64 = float64(0.0)
        _ = v8
        _ = v9
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v1)
        }
        v10 := GenS1{F0: v1, F1: 1, F2: 1}
        v11 := GenS3{F0: -103, F1: 797}
        _ = v10
        _ = v11
    }
    v12 := GenS4{F0: 0, F1: v3}
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS4{F0: fi13, F1: 442}
        if false {
            v3 -= v3
        } else {
            var v15 uint32 = v3
            _ = v15
        }
        v1 += v1
        _ = fi13
        _ = v14
    }
    var v16 int16 = -348
    var v17 int = -1
    v17 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v18 float32 = 0.0
    var v19 uint16 = 738
    switch "hello" {
    case "":
        var v20 float32 = (1e10 * 0.0)
        _ = v20
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v2 = append(v2, v3)
    }
    v18 -= (1e10 - 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
