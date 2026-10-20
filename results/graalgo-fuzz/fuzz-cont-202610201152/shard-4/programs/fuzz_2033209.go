package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0(p0 bool, p1 int32) uint {
    switch 32768 {
    case 439:
        for fi5 := 0; fi5 < 2; fi5++ {
            if (127 <= -6) {
                switch "hello" {
                case "foo bar":
                    var v6 int32 = (p1 + p1)
                    _ = v6
                case "a":
                    v7 := GenS1{F0: 65535, F1: 4294967295}
                    _ = v7
                case "\n\t":
                    fmt.Println(p1, p1, fi5)
                default:
                    var v8 int8 = -1
                    _ = v8
                }
                p1 = (p1 | -2147483648)
            }
            _ = fi5
        }
    case 659:
        var v9 uint16 = 32767
        _ = v9
    case 1:
        fmt.Println(p0, p1)
    default:
        v10 := []int{-648, 9223372036854775807, 2147483647}
        _ = v10
    }
    if true {
        v11 := GenS2{F0: 4294967295, F1: 9}
        switch 127 {
        case -9223372036854775808:
            p0 = true
        case -1:
            v12 := GenS2{F0: 0, F1: 0}
            _ = v12
        case 4294967295:
            if (431 == 762) {
                var v13 int16 = 341
                v13 += v13
                _ = v13
            }
        default:
            p0 = true
        }
        _ = v11
    }
    v14 := []int{127, 4294967295, 0, -102}
    _ = p0
    _ = p1
    _ = v14
    return 859
}
func genFn1(p0 uint64, p1 int16, p2 int32) int {
    if true {
        fmt.Println(p0, p1)
    }
    p0 = (^(p0))
    switch ("foo bar" + ("foo bar" + "hello")) {
    case "":
        fmt.Println(p1, p1, p2)
    case "foo bar":
        var v15 uint32 = 716
        _ = v15
    case "a":
        p1 = p1
    default:
        v16 := []int{42, -629, -1}
        _ = v16
    }
    v17 := GenS2{F0: p0, F1: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -527
}

func main() {
    var v1 float32 = -2.25
    var v2 int = 96
    var v3 int64 = 1
    v4 := []uint32{797, 2147483648, 42}
    var v18 int16 = 1000
    v1 -= v1
    v19 := GenS2{F0: 438, F1: 82}
    fmt.Println(v3, v3)
    if false {
        switch "hello" {
        case "hello":
            v4 = append(v4, 4294967295)
        default:
            v3 = v3
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            v4 = append(v4, 2147483647)
            fmt.Println(v1)
            fmt.Println(fi20, v18)
            _ = fi20
        }
        v21 := GenS3{F0: 18446744073709551615, F1: 760, F2: "foo bar"}
        _ = v21
    } else {
        v4 = append(v4, 241)
    }
    v1 = v1
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 int = fi22
        v1 = 0.1
        _ = fi22
        _ = v23
    }
    var v24 uint8 = 255
    var v25 int8 = 127
    v3 -= v3
    if false {
        _ = len(v4)
    } else {
        _ = len(v4)
    }
    v26 := GenS1{F0: 65535, F1: 1}
    v27 := GenS4{F0: 786, F1: (v25 << ((0) & 63)), F2: 2}
    v18 += (^(v18))
    var v28 uint16 = 255
    var v29 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
