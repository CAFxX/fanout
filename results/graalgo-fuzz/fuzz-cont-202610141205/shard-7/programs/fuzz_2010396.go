package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int; F2 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16) string {
    var v4 int = 1
    fmt.Println(v4)
    if true {
        var v5 uint64 = 2
        p0 += p0
        v6 := GenS1{F0: 255, F1: uint16(v5), F2: int64(v4)}
        _ = v5
        _ = v6
    }
    var v7 string = "foo bar"
    _ = p0
    _ = v4
    _ = v7
    return v7
}
func genFn1(p0 uint32, p1 int8) uint8 {
    v8 := GenS2{F0: 42}
    switch ("" + (((("a" + "a") + "") + "a") + "foo bar")) {
    case "":
        fmt.Println(p0, p0)
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            p1 = (p1 / (((p1 % (((p1 << ((1) & 63))) | 1))) | 1))
            v10 := GenS1{F0: 42, F1: 255, F2: -81}
            p0 += (p0 / ((p0) | 1))
            _ = fi9
            _ = v10
        }
    default:
        var v11 uint64 = 987
        _ = v11
    }
    v12 := []int{}
    _ = len(v12)
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    return 60
}

func main() {
    v1 := []string{(("a" + "hello") + "\n\t"), "foo bar", "foo bar", "hello"}
    v2 := make(map[uint32]bool)
    var v3 int32 = 2147483647
    v3 = v3
    switch "\n\t" {
    case "foo bar":
        var v13 uint32 = 867
        _ = v13
    case "":
        _ = len(v1)
    case "a":
        switch -9223372036854775808 {
        case 1:
            v14 := GenS3{F0: 42, F1: 0}
            _ = v14
        default:
            switch (("foo bar" + "") + "a") {
            case "":
                delete(v2, 2147483648)
            default:
                v15 := v2[2]
                _ = v15
            }
        }
    }
    v3 += (v3 * (v3 - (v3 ^ (v3 & 1))))
    v16 := GenS4{F0: 444, F1: 42}
    switch "a" {
    case "a":
        v17 := GenS3{F0: 0, F1: -128}
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(fi18)
            v19 := GenS5{F0: 1, F1: fi18, F2: 0}
            _ = fi18
            _ = v19
        }
    default:
        v1 = append(v1, ("\n\t" + "\n\t"))
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS6{F0: 1}
        var v22 int = -658
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int8 = -36
    fmt.Println(v3, v23)
    switch -2147483648 {
    case 1:
        var v24 int8 = (v23 << ((1) & 63))
        _ = v24
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -642:
        v3 = -1
    default:
        v23 += v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
