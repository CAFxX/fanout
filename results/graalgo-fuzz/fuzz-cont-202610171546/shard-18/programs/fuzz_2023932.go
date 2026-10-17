package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int16 }
func genFn0(p0 int, p1 int64) uint64 {
    fmt.Println(p1, p0)
    var v7 int64 = (p1 & 180)
    v8 := []int{117}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 42
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v5 uint32 = 0
    var v6 uint16 = 260
    switch "a" {
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 bool = ((false && (((false && ((v6 >> ((377) & 63)) <= (v6 + v6))) && true) && false)) || true)
            _ = fi9
            _ = v10
        }
    case "foo bar":
        v6 -= (v6 / ((v6) | 1))
    default:
        v11 := GenS1{F0: 278, F1: -232, F2: -54}
        _ = v11
    }
    var v12 uint = 42
    fmt.Println(v2)
    var v13 uint32 = 894
    if ((v5 % ((1) | 1)) <= v13) {
        v14 := GenS2{F0: 42, F1: 216, F2: 660}
        _ = v14
    }
    var v15 uint8 = uint8(v12)
    var v16 int32 = 2147483647
    if true {
        v5 += 2147483647
        var v17 uint32 = 205
        switch "foo bar" {
        case "hello":
            v1 = append(v1, (^(v17)))
        }
        _ = v17
    } else {
        var v18 uint8 = v15
        for fi19 := 0; fi19 < 20; fi19++ {
            if (0.0 <= 1e10) {
                v20 := GenS3{F0: (v16 >> ((1) & 63))}
                _ = v20
            } else {
                v21 := GenS4{F0: uint16(v15), F1: v18, F2: 1000}
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 float64 = 1.5
                    var v24 uint8 = uint8(v6)
                    var v25 int64 = 9223372036854775807
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                fmt.Println(v2)
                _ = v21
            }
            _ = fi19
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 string = "hello"
            _ = fi26
            _ = v27
        }
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
