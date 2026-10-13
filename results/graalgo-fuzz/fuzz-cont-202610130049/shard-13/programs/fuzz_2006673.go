package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int64; F1 uint; F2 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0() int8 {
    var v6 uint16 = 2
    switch 32768 {
    case 15:
        fmt.Println(v6)
    case 0:
        var v7 uint64 = 0
        _ = v7
    default:
        switch 127 {
        case 9223372036854775807:
            v8 := GenS1{F0: 42, F1: 1, F2: 42}
            _ = v8
        }
    }
    fmt.Println(v6)
    _ = v6
    return -106
}
func genFn1(p0 uint16) string {
    fmt.Println(p0)
    switch "hello" {
    case "hello":
        if ((1 == 1) || false) {
            p0 = p0
        }
    case "\n\t":
        fmt.Println(p0)
    case "a":
        if ("a" != (("foo bar" + "hello") + "hello")) {
            v9 := GenS2{F0: 127, F1: 0, F2: ("\n\t" + "a")}
            _ = v9
        } else {
            v10 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807, F2: -1}
            var v11 bool = true
            _ = v10
            _ = v11
        }
    }
    p0 = p0
    v12 := GenS3{F0: 1, F1: 2, F2: 173}
    v13 := GenS4{F0: -259, F1: 478}
    _ = p0
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 int8 = 80
    v2 := make(map[int]uint32)
    var v3 string = "\n\t"
    var v4 uint16 = 65535
    v5 := make(map[uint32]uint32)
    v3 = "hello"
    v14 := GenS5{F0: 781, F1: 128}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint8 = 0
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v3 = ("\n\t" + v3)
    switch "" {
    case "foo bar":
        var v18 int16 = -56
        _ = v18
    case "":
        if (!false) {
            _ = len(v5)
            fmt.Println(v4, v1, v1)
        }
    case "\n\t":
        v4 -= 2
    default:
        for fi19 := 0; fi19 < 12; fi19++ {
            if false {
                v1 += (v1 / (((v1 / (((-(v1))) | 1))) | 1))
                v1 -= v1
            }
            _ = fi19
        }
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint16 = 2
        _ = fi20
        _ = v21
    }
    v4 += (v4 + (v4 & 187))
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int8 = genFn0()
        _ = fi22
        _ = v23
    }
    v24 := []int{-117, -468, int(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
