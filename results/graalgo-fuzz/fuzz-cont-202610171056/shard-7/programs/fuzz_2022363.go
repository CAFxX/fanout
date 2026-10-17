package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64, p1 int8, p2 int32) int {
    v5 := []int{9223372036854775807, 695, -379}
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 = append(v5, (fi6 * fi6))
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -929
}
func genFn1(p0 bool) string {
    var v7 int64 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        switch (^(fi8)) {
        case 965:
            v9 := GenS1{F0: 1}
            _ = v9
        case 4294967295:
            v10 := GenS2{F0: true, F1: 42, F2: 0}
            _ = v10
        default:
            var v11 float32 = 0.0
            _ = v11
        }
        var v12 int = genFn0(2, -95, 739)
        _ = fi8
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = v13
    _ = v7
    return "a"
}

func main() {
    v1 := []string{("a" + "a"), ""}
    var v2 int = -2147483648
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v14 uint64 = 0
    v2 = v2
    if (v14 > 42) {
        fmt.Println(v2)
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 = append(v1, "foo bar")
            _ = fi15
        }
        _ = len(v1)
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int = (fi16 * fi16)
        v4[(fi16 ^ (^(v17)))] = 4294967295
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v14)
    if true {
        v18 := GenS3{F0: 2147483648}
        _ = v18
    }
    if ((v2 <= 32767) || false) {
        v14 = (v14 << ((343) & 63))
        v19 := GenS4{F0: true}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v19
    }
    if false {
        v20 := GenS3{F0: 764}
        v2 = 127
        _ = v20
    } else {
        v21 := GenS5{F0: 2147483647, F1: 100}
        if true {
            var v22 bool = (1 == 1)
            if (uint8(v14) == 85) {
                var v23 uint = 18446744073709551615
                v24 := GenS5{F0: 42, F1: -1}
                switch "foo bar" {
                case "\n\t":
                    v14 -= (v14 / (((^(v14))) | 1))
                case "hello":
                    v2 = (v2 | v2)
                }
                _ = v23
                _ = v24
            } else {
                var v25 float32 = (0.1 / 1.5)
                if ((!v22) || (!(!(((true || v22) && ((1.5 * float64(3.14159)) > -2.25)) || true)))) {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 uint32 = 715
                        _ = fi26
                        _ = v27
                    }
                } else {
                    var v28 int = v2
                    v2 = genFn0((^(v14)), -1, -1)
                    _ = v28
                }
                var v29 int32 = -2147483648
                _ = v25
                _ = v29
            }
            v2 = v2
            _ = v22
        }
        v3[(-(v2))] = 4294967295
        _ = v21
    }
    v4[v2] = 180
    v30 := GenS3{F0: 291}
    v2 = (v2 ^ (^(v2)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
