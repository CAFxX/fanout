package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
func genFn0(p0 string, p1 uint16) uint8 {
    p0 = p0
    var v5 string = "\n\t"
    var v6 uint = 255
    v6 = (^(v6))
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 64
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 2
    v3 := []int{4294967295, 9223372036854775807, -2147483648, 128}
    var v4 string = "a"
    var v8 int = -963
    v4 = "hello"
    if false {
        fmt.Println(v2, v2)
    } else {
        if true {
            switch v4 {
            case "a":
                v2 += (v2 / (((^(v2))) | 1))
            case "foo bar":
                if (!(!true)) {
                    var v9 bool = (!(32767 <= -75))
                    _ = len(v1)
                    _ = v9
                } else {
                    var v10 int64 = 9223372036854775807
                    var v11 float32 = 0.1
                    _ = v10
                    _ = v11
                }
            case "\n\t":
                v12 := GenS1{F0: 65535, F1: 0, F2: 18446744073709551615}
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int = 128
        v15 := GenS1{F0: 0, F1: v8, F2: 18446744073709551615}
        _ = fi13
        _ = v14
        _ = v15
    }
    v8 += (-(v8))
    switch v8 {
    case 768:
        switch (v4 + v4) {
        case "hello":
            v4 = ("a" + v4)
        case "foo bar":
            v4 = v4
        case "\n\t":
            v8 = (v8 / (((v8 & v8)) | 1))
        default:
            v4 = v4
        }
    default:
        _ = len(v1)
    }
    v2 += (v2 << ((2) & 63))
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: 413, F1: -866, F2: 0}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 1, F1: v8, F2: 255}
    _ = len(v1)
    v8 -= (v8 | v8)
    if (51 == 42) {
        v19 := GenS2{F0: 2, F1: -1, F2: 2147483647}
        for fi20 := 0; fi20 < 13; fi20++ {
            _ = len(v1)
            _ = fi20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v4 = v4
        _ = fi21
    }
    v8 -= (v8 - (v8 / (((v8 | v8)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
