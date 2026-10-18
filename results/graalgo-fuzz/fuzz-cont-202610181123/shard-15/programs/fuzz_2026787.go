package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16, p1 int, p2 uint32) int32 {
    fmt.Println(p0, p0)
    if false {
        fmt.Println(p2, p0, p2)
        switch ("" + "\n\t") {
        case "\n\t":
            v4 := []int{(p1 << ((700) & 63)), (p1 ^ p1), 42}
            _ = v4
        }
        var v5 bool = false
        _ = v5
    } else {
        v6 := []int{}
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            fmt.Println(p2)
            var v7 uint = 4294967295
            _ = v7
        }
        v6 = append(v6, (p1 >> ((255) & 63)))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return -99
}

func main() {
    v1 := []int{}
    v2 := []uint32{1, 65535}
    v3 := []int{-952, 127}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 850, F1: -9223372036854775808}
        var v10 float64 = 1e10
        fmt.Println(fi8, v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    if (1e10 <= 0.1) {
        v11 := GenS2{F0: 42}
        _ = v11
    } else {
        var v12 uint32 = 2
        fmt.Println(v12)
        _ = v12
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    var v13 uint32 = 2147483647
    switch 0 {
    case -1:
        for fi14 := 0; fi14 < 17; fi14++ {
            if (1 != -906) {
                var v15 uint64 = 146
                _ = v15
            } else {
                v13 = v13
                if len(v3) > 0 {
                    v3[0] = -663
                }
                v13 = (v13 | 333)
            }
            v13 += v13
            _ = fi14
        }
    case 32768:
        v13 = v13
    case -162:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "\n\t" {
    case "\n\t":
        v13 -= 42
    case "foo bar":
        var v16 uint64 = 874
        _ = v16
    default:
        v17 := GenS1{F0: 32767, F1: 42}
        _ = v17
    }
    v18 := GenS2{F0: 9223372036854775807}
    v13 += (v13 >> ((825) & 63))
    v13 = (v13 & (v13 * v13))
    var v19 int32 = 2147483647
    fmt.Println(v13, v13)
    v20 := GenS3{F0: 1, F1: int64(v19), F2: true}
    v21 := GenS4{F0: 2, F1: "a"}
    switch "foo bar" {
    case "":
        switch "a" {
        case "":
            v22 := GenS5{F0: 4294967295}
            _ = v22
        case "\n\t":
            v19 = v19
        case "a":
            if true {
                if true {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 int8 = -128
                        v25 := GenS3{F0: 100, F1: 9223372036854775807, F2: (664 > 255)}
                        if len(v3) > 0 {
                            v3[0] = fi23
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint32 = (v13 ^ v13)
                    v1 = append(v1, -1)
                    _ = v26
                } else {
                    v13 += (v13 - v13)
                    v19 = (-(v19))
                    fmt.Println(v19, v19)
                }
            }
        }
    case "foo bar":
        fmt.Println(v19, v13)
    default:
        v27 := GenS6{F0: 2, F1: 2}
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
