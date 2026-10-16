package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 uint8, p1 int64, p2 uint64) uint {
    var v7 string = "hello"
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 42, F1: uint32(p0)}
            if ((false || true) || (66 != -84)) {
                v10 := []int{(^(fi8)), fi8, fi8}
                switch v7 {
                case "":
                    v11 := GenS1{F0: p2, F1: 65535}
                    _ = v11
                case "\n\t":
                    fmt.Println(p1)
                case "a":
                    var v12 int8 = 1
                    _ = v12
                default:
                    v7 = v7
                }
                _ = v10
            } else {
                v13 := GenS2{F0: 1, F1: (((v7 + v7) + v7) + (("\n\t" + ((((v7 + v7) + v7) + "a") + v7)) + "foo bar")), F2: 1}
                _ = v13
            }
            var v14 int16 = 1000
            _ = fi8
            _ = v14
            _ = v9
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(p0, p2, p0)
            if (false || (-128 == 0)) {
                if (uint32(p2) >= 1) {
                    p2 = (p2 << ((0) & 63))
                }
            } else {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := []int{(-(fi15)), (-(fi16)), fi16}
                    if (253 <= 2) {
                        p2 -= (p2 ^ (p2 * p2))
                    }
                    if len(v17) > 0 {
                        v17[0] = (fi16 / ((32768) | 1))
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
    }
    v18 := GenS1{F0: p2, F1: 4294967295}
    p0 += 42
    var v19 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v7
    return 255
}
func genFn1(p0 int16) int8 {
    switch -1 {
    case 1:
        fmt.Println(p0)
    case 9223372036854775807:
        var v20 bool = true
        _ = v20
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]bool)
    var v3 int64 = 0
    var v4 int = v1
    v5 := []int{(v1 << ((1) & 63))}
    var v6 int16 = 1000
    if true {
        v3 = (v3 - v3)
    }
    var v21 int16 = v6
    var v22 uint16 = 132
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 int32 = 0
        v3 = (v3 << ((0) & 63))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v5 = append(v5, -623)
        switch (fi25 & (v1 << ((139) & 63))) {
        case 1:
            if true {
                v26 := v2[4294967295]
                _ = v26
            } else {
                fmt.Println(v6)
                switch "" {
                case "a":
                    var v27 int16 = v6
                    _ = v27
                case "\n\t":
                    _ = len(v2)
                case "hello":
                    fmt.Println(v21, v21)
                default:
                    v3 += v3
                }
                v21 = -634
            }
        case 0:
            v22 = v22
        default:
            delete(v2, 1)
        }
        _ = fi25
    }
    v6 -= v21
    switch (v1 | v1) {
    case 0:
        v3 += 0
    case 1:
        if (!false) {
            _ = len(v5)
            v3 = -1
        }
    case -1:
        var v28 bool = false
        _ = v28
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
