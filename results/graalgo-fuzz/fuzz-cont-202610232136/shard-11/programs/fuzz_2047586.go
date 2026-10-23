package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 uint16, p1 int32, p2 int8) int64 {
    var v7 uint16 = (p0 / ((526) | 1))
    if ((0.1 < ((3.14159 * float64(1.5)) * (0.1 * (((0.0 * 0.0) / -2.25) + (0.0 + 0.1))))) || (42 < (^(p1)))) {
        p0 = (^(v7))
        v8 := []int{}
        v7 = p0
        _ = v8
    } else {
        fmt.Println(p1, p2, p0)
        var v9 float32 = 3.14159
        _ = v9
    }
    p1 += (-(p1))
    fmt.Println(p1)
    if false {
        v10 := GenS1{F0: 1, F1: 926}
        v11 := []int{127, -1, -1, 4294967295}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -710
}
func genFn1(p0 int64, p1 uint16, p2 uint64) uint32 {
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := []int{-171, 981}
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS2{F0: (fi12 & fi12)}
            switch "hello" {
            case "foo bar":
                if false {
                    fmt.Println(fi14, fi12)
                    fmt.Println(fi12)
                } else {
                    switch "hello" {
                    case "hello":
                        var v16 uint32 = 2147483647
                        _ = v16
                    case "foo bar":
                        var v17 int8 = -1
                        _ = v17
                    }
                    p0 = p0
                }
            case "hello":
                if true {
                    p2 += (p2 ^ 42)
                } else {
                    p2 -= 621
                    v18 := GenS3{F0: (p0 + (p0 << ((18446744073709551615) & 63))), F1: uint32(p1), F2: ("hello" + "\n\t")}
                    var v19 uint = 4294967295
                    _ = v18
                    _ = v19
                }
            default:
                var v20 uint16 = p1
                _ = v20
            }
            fmt.Println(fi14, fi12)
            var v21 uint32 = 0
            _ = fi14
            _ = v15
            _ = v21
        }
        fmt.Println(fi12, p0)
        _ = fi12
        _ = v13
    }
    v22 := []int{164}
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 2147483647
}

func main() {
    var v1 bool = false
    v2 := []int{}
    v3 := make(map[string]bool)
    var v4 uint64 = 4294967295
    var v5 bool = v1
    var v6 int8 = -123
    var v23 uint64 = 2
    v4 -= (v4 - 73)
    var v24 int64 = 9223372036854775807
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            v1 = true
            _ = fi26
        }
        var v27 int8 = v6
        _ = fi25
        _ = v27
    }
    v4 += (v23 ^ (^(v4)))
    var v28 uint64 = v4
    var v29 int8 = 127
    v30 := GenS1{F0: 92, F1: 65535}
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v4, v1)
    }
    switch "a" {
    case "a":
        fmt.Println(v29, v29, v5)
    case "hello":
        fmt.Println(v28, v23)
    case "\n\t":
        _ = len(v2)
    default:
        if ((v23 - v28) < (^(v23))) {
            v1 = (255 <= 244)
        } else {
            var v31 bool = (65 != 0)
            var v32 int32 = 1
            _ = v31
            _ = v32
        }
    }
    v5 = (42 < 32767)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
