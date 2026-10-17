package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 uint8) bool {
    var v4 int64 = -467
    v5 := GenS1{F0: 42}
    if true {
        v4 = int64(p0)
        v6 := []int{}
        v7 := GenS1{F0: 2147483648}
        _ = v6
        _ = v7
    } else {
        if true {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{(fi8 / ((-111) | 1)), (fi8 & (fi8 | (fi8 << ((65535) & 63)))), fi8, (fi8 << ((4294967295) & 63))}
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: -186}
            v4 += v4
            _ = v10
        }
    }
    v4 -= v4
    _ = p0
    _ = v4
    _ = v5
    return false
}

func main() {
    v1 := []uint32{954, 2}
    var v2 float32 = float32((3.14159 + 1e10))
    var v3 int = 128
    var v11 uint16 = 32767
    var v12 int64 = 42
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v2, v3)
        _ = fi13
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        switch "hello" {
        case "hello":
            v12 += -9223372036854775808
        case "\n\t":
            v3 = 42
        default:
            v2 -= float32(-2.25)
        }
    default:
        fmt.Println(v11)
    }
    if (!(v12 >= (^(v12)))) {
        _ = len(v1)
    } else {
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(v11, v3)
            switch 572 {
            case 42:
                v12 = v12
            case 94:
                v11 -= v11
            default:
                var v15 int32 = 2147483647
                _ = v15
            }
            _ = fi14
        }
        var v16 int32 = -500
        _ = v16
    }
    switch "a" {
    case "hello":
        if ((v2 / v2) >= v2) {
            fmt.Println(v2, v2)
            var v17 uint = 925
            _ = v17
        }
    case "foo bar":
        var v18 int8 = 100
        _ = v18
    default:
        var v19 int16 = 32767
        _ = v19
    }
    v12 = v12
    v2 += ((1e10 + v2) / (1.5 / 1e10))
    v3 = 4294967295
    if true {
        v20 := GenS3{F0: -1, F1: "hello"}
        v12 = (^(v12))
        _ = v20
    }
    v21 := GenS4{F0: 492, F1: 764}
    if (true || false) {
        v12 -= v12
        var v22 string = ("" + "hello")
        _ = v22
    }
    v3 += 265
    if len(v1) > 0 {
        v1[0] = 724
    }
    if true {
        for fi23 := 0; fi23 < 14; fi23++ {
            switch "foo bar" {
            case "":
                var v24 int16 = -882
                _ = v24
            case "a":
                _ = len(v1)
            case "foo bar":
                v25 := GenS1{F0: 1}
                _ = v25
            }
            fmt.Println(fi23, v12)
            _ = fi23
        }
        fmt.Println(v2)
        fmt.Println(v12, v11, v12)
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
}
