package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
type GenS4 struct { F0 bool }
func genFn0(p0 int32, p1 uint, p2 int8) string {
    v5 := []int{}
    v6 := GenS1{F0: 459}
    v7 := GenS1{F0: 0}
    if true {
        fmt.Println(p0, p0)
        var v8 uint = p1
        fmt.Println(p0)
        _ = v8
    }
    var v9 int32 = int32(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return (("\n\t" + ("a" + "a")) + "a")
}
func genFn1(p0 int64) int32 {
    switch ("\n\t" + "hello") {
    case "\n\t":
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{(fi10 & (fi10 / (((fi10 << ((42) & 63))) | 1))), (^(fi10)), (fi10 ^ fi10), fi10}
            _ = fi10
            _ = v11
        }
    case "hello":
        var v12 uint8 = 127
        _ = v12
    case "foo bar":
        var v13 uint64 = 902
        _ = v13
    default:
        var v14 uint = 494
        _ = v14
    }
    if true {
        p0 = (-(p0))
    }
    if true {
        var v15 uint = 255
        _ = v15
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 uint64 = 0
    var v2 float64 = -2.25
    var v3 int16 = -1
    v4 := make(map[uint32]bool)
    fmt.Println(v1, v2, v2)
    v2 -= 0.1
    switch 4294967295 {
    case -859:
        v3 = 285
    case -106:
        var v16 uint8 = 55
        _ = v16
    default:
        if (255 >= 32767) {
            v17 := v4[297]
            v18 := []int{32768, -164, 9223372036854775807, -659}
            _ = v17
            _ = v18
        } else {
            switch "\n\t" {
            case "foo bar":
                v19 := GenS2{F0: (1 >= 100), F1: 9223372036854775807}
                _ = v19
            case "\n\t":
                v3 = v3
            case "":
                if (522 > 65535) {
                    var v20 int64 = -931
                    v1 -= v1
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int32 = 0
                        v2 = 1e10
                        _ = fi21
                        _ = v22
                    }
                    v3 = (-(v3))
                }
            default:
                v23 := GenS3{F0: 18446744073709551615, F1: true, F2: ""}
                _ = v23
            }
        }
    }
    v2 -= v2
    v3 = (v3 ^ v3)
    var v24 bool = ((1e10 - (float64(1e10) + v2)) == v2)
    var v25 int32 = genFn1(-821)
    v2 += v2
    v26 := GenS4{F0: (true && (0 > 65535))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
