package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16) int64 {
    v6 := GenS1{F0: -750, F1: false, F2: p0}
    switch -9223372036854775808 {
    case -9223372036854775808:
        p0 = (^(p0))
    case -220:
        var v7 uint16 = (p0 - (p0 % (((^(p0))) | 1)))
        _ = v7
    case 1:
        var v8 int16 = 0
        _ = v8
    default:
        var v9 float32 = 3.14159
        _ = v9
    }
    if false {
        v10 := []int{-2147483648, 9223372036854775807, 1}
        var v11 int8 = 1
        p0 -= p0
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                p0 -= p0
            }
            _ = fi12
        }
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    _ = p0
    _ = v6
    return 134
}
func genFn1(p0 uint8, p1 uint, p2 uint16) int8 {
    var v14 uint = 1
    v14 = 4294967295
    p0 = (p0 * (p0 * uint8(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -128
}

func main() {
    var v1 int = 221
    var v2 uint8 = 0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2}
    v5 := make(map[uint32]string)
    switch (("a" + "foo bar") + "a") {
    case "a":
        if false {
            switch (^(v1)) {
            case -9223372036854775808:
                _ = len(v5)
            case 42:
                switch v1 {
                case 2147483647:
                    var v15 float64 = float64(3.14159)
                    _ = v15
                case 4294967295:
                    if true {
                        v1 = 32767
                    }
                case 304:
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 uint16 = 65535
                        var v18 float64 = (-2.25 / (1.5 + 1e10))
                        v19 := GenS2{F0: 0}
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
            default:
                v20 := GenS2{F0: 18446744073709551615}
                _ = v20
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 bool = true
                fmt.Println(v1, v2)
                _ = fi21
                _ = v22
            }
        } else {
            v2 += v2
            var v23 bool = false
            var v24 int64 = 421
            _ = v23
            _ = v24
        }
    case "hello":
        v25 := GenS3{F0: v1}
        _ = v25
    }
    v26 := GenS4{F0: 0, F1: (v2 >> ((174) & 63))}
    if (2 >= 31) {
        for fi27 := 0; fi27 < 18; fi27++ {
            v1 = v1
            _ = fi27
        }
        _ = len(v3)
        var v28 int64 = 936
        _ = v28
    } else {
        var v29 float32 = 1.5
        v1 = (v1 + 378)
        _ = v29
    }
    v1 += (v1 << ((2) & 63))
    v1 = v1
    v30 := GenS5{F0: 2147483647}
    v2 -= 42
    v4 = append(v4, 360)
    fmt.Println(v2)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
