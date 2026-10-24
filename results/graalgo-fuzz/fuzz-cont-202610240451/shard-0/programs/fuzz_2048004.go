package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 uint, p1 int16) string {
    var v4 int = 42
    fmt.Println(p1, v4, p1)
    if (1 < 65535) {
        var v5 float64 = ((1e10 + (float64(0.0) * 1e10)) / float64(float32((((1.5 - 1e10) / 0.1) - 0.0))))
        v6 := GenS1{F0: (-(p1)), F1: -2147483648, F2: (false && false)}
        v7 := []int{}
        _ = v5
        _ = v6
        _ = v7
    } else {
        if (0.1 == 0.0) {
            for fi8 := 0; fi8 < 11; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    switch "hello" {
                    case "":
                        fmt.Println(fi8, fi8, v4)
                    case "\n\t":
                        p0 += (^(p0))
                    default:
                        var v10 uint8 = 42
                        _ = v10
                    }
                    fmt.Println(v4)
                    _ = fi9
                }
                p0 += p0
                _ = fi8
            }
        }
    }
    v11 := []int{(v4 << ((600) & 63)), v4, (v4 - v4)}
    var v12 int = (v4 ^ (v4 ^ (v4 | (v4 | v4))))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v4
    return ""
}
func genFn1(p0 int) int64 {
    if false {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint = 744
            var v15 int64 = 9223372036854775807
            _ = fi13
            _ = v14
            _ = v15
        }
        p0 -= 69
    } else {
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := []int{-1, p0, 127, (p0 & (fi16 ^ fi16))}
            _ = fi16
            _ = v17
        }
        v18 := []int{p0, (p0 << ((36) & 63)), p0}
        v19 := GenS2{F0: -1, F1: -1}
        _ = v18
        _ = v19
    }
    var v20 uint = 249
    v21 := GenS2{F0: -790, F1: -1}
    _ = p0
    _ = v20
    _ = v21
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v22 uint16 = 65535
    v23 := v2[2147483648]
    if true {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS2{F0: -2147483648, F1: 700}
            _ = fi24
            _ = v25
        }
    }
    v22 += (^(v22))
    var v26 uint32 = (^(v23))
    v26 += 893
    v26 -= (v26 << ((236) & 63))
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 int16 = 1
        _ = fi27
        _ = v28
    }
    var v29 int32 = -407
    switch "hello" {
    case "\n\t":
        if (!(737 == 223)) {
            v30 := []int{-157, 2147483647, -846, 519}
            _ = v30
        } else {
            v26 = v23
            v31 := v1[398]
            if (false || true) {
                for fi32 := 0; fi32 < 16; fi32++ {
                    switch (v31 << ((83) & 63)) {
                    case 4294967295:
                        v23 = v23
                    case -9223372036854775808:
                        fmt.Println(v31, v22, v26)
                    }
                    fmt.Println(fi32, fi32)
                    _ = fi32
                }
            }
            _ = v31
        }
    default:
        if false {
            fmt.Println(v23, v22)
            v33 := []int{380, 128, 2147483647}
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v33
        }
    }
    var v34 uint64 = 4294967295
    _ = len(v3)
    v35 := []int{910, -1, 0, -737}
    v36 := GenS2{F0: 1, F1: v29}
    v37 := GenS1{F0: 0, F1: (^(v29)), F2: false}
    var v38 uint8 = 42
    fmt.Println(v34, v34, v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
