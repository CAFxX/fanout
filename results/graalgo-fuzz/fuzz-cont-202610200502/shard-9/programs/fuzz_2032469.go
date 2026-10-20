package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint32) int64 {
    if (683 <= 352) {
        v6 := GenS1{F0: 32767, F1: 18446744073709551615}
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 -= p0
            switch fi7 {
            case -9223372036854775808:
                fmt.Println(p0)
            default:
                if true {
                    p0 += (p0 & 114)
                    var v8 uint32 = (p0 / ((840) | 1))
                    fmt.Println(p0, fi7, p0)
                    _ = v8
                } else {
                    var v9 uint32 = p0
                    if false {
                        v10 := []int{-313}
                        _ = v10
                    }
                    _ = v9
                }
            }
            v11 := GenS2{F0: 1, F1: (-(fi7))}
            _ = fi7
            _ = v11
        }
        _ = v6
    } else {
        fmt.Println(p0)
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 0, F1: 42}
        _ = fi12
        _ = v13
    }
    _ = p0
    return 42
}
func genFn1(p0 uint8, p1 uint32, p2 int16) int16 {
    switch "foo bar" {
    case "":
        p1 += 2
    case "a":
        p0 = (p0 >> ((181) & 63))
    default:
        var v14 uint32 = 1
        _ = v14
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 797
}

func main() {
    var v1 int32 = -491
    v2 := make(map[int]bool)
    v3 := []int{-680}
    var v4 uint8 = 1
    v5 := make(map[int]int)
    _ = len(v2)
    var v15 int8 = 0
    switch "hello" {
    case "\n\t":
        if (v4 == (v4 - v4)) {
            var v16 int8 = v15
            _ = v16
        }
    case "hello":
        v15 = (^(v15))
    case "foo bar":
        var v17 int = 32768
        _ = v17
    default:
        v1 -= (-(v1))
    }
    if false {
        v18 := GenS3{F0: 806}
        _ = v18
    }
    switch "foo bar" {
    case "a":
        var v19 int8 = v15
        _ = v19
    }
    v15 -= -128
    if false {
        v20 := GenS2{F0: 999, F1: 1}
        fmt.Println(v4)
        _ = v20
    } else {
        var v21 uint16 = 32767
        v21 = v21
        v15 += (^(v15))
        _ = v21
    }
    v3 = append(v3, 0)
    var v22 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
