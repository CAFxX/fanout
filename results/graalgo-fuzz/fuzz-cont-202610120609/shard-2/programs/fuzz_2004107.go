package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint32; F1 uint32; F2 int16 }
func genFn0(p0 int32, p1 string) int16 {
    switch "\n\t" {
    case "a":
        var v5 int32 = (p0 | (p0 & (p0 ^ p0)))
        _ = v5
    case "\n\t":
        v6 := GenS1{F0: int64(p0), F1: p1, F2: 654}
        _ = v6
    }
    p1 = "a"
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint16, p1 int8, p2 uint) bool {
    switch (("" + ("hello" + ("foo bar" + ("" + "\n\t")))) + "a") {
    case "a":
        var v7 float32 = float32(0.0)
        _ = v7
    default:
        var v8 int32 = 2147483647
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 278, F1: "a", F2: 358}
        if true {
            var v11 int64 = 0
            fmt.Println(p2, v11)
            _ = v11
        }
        v12 := GenS1{F0: -684, F1: ("hello" + (("" + "\n\t") + "\n\t")), F2: 505}
        _ = fi9
        _ = v10
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int8 = 0
    v2 := []int{}
    var v3 int64 = -855
    var v4 int16 = -70
    v3 = v3
    _ = len(v2)
    for fi13 := 0; fi13 < 3; fi13++ {
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = fi13
    }
    if (true && true) {
        fmt.Println(v1)
        switch 42 {
        case 669:
            fmt.Println(v3)
        case 9223372036854775807:
            var v14 string = "\n\t"
            _ = v14
        case -9223372036854775808:
            var v15 uint = 42
            _ = v15
        default:
            if (634 >= 2147483647) {
                var v16 float64 = (0.0 - (-2.25 / -2.25))
                _ = v16
            } else {
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (!((v3 ^ v3) != (-(v3)))) {
                        v3 = (-(v3))
                    }
                    _ = fi17
                }
            }
        }
        if true {
            fmt.Println(v1)
        } else {
            v18 := GenS2{F0: false}
            v4 = (v4 & v4)
            _ = v18
        }
    }
    v4 = v4
    if len(v2) > 0 {
        v2[0] = 32768
    }
    fmt.Println(v3)
    fmt.Println(v1)
    var v19 int64 = v3
    v3 += 0
    v3 -= (^(v3))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "hello":
        for fi20 := 0; fi20 < 13; fi20++ {
            if ("foo bar" >= ("a" + "hello")) {
                v21 := GenS3{F0: v4, F1: fi20}
                v4 = v4
                v3 = v19
                _ = v21
            }
            v2 = append(v2, (fi20 % ((fi20) | 1)))
            _ = len(v2)
            _ = fi20
        }
    case "\n\t":
        var v22 uint64 = 762
        _ = v22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    v23 := GenS4{F0: 1, F1: 1, F2: 1}
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS3{F0: 439, F1: fi24}
        v4 -= 0
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
