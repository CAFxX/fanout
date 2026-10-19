package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 uint8, p1 string, p2 bool) uint {
    p2 = p2
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: 42, F1: ((("\n\t" + ("\n\t" + p1)) + "a") + p1), F2: 0}
        p1 = "a"
        _ = fi5
        _ = v6
    }
    p2 = (p0 < 1)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 uint8, p1 int64, p2 string) int8 {
    var v7 int16 = 2
    p1 += -997
    if (42 < 42) {
        if true {
            if (!(-564 < p1)) {
                var v8 int16 = v7
                p0 = (p0 << ((255) & 63))
                _ = v8
            } else {
                var v9 uint64 = 18446744073709551615
                switch -2147483648 {
                case 128:
                    v10 := GenS1{F0: (^(v7)), F1: "\n\t", F2: 65535}
                    _ = v10
                }
                _ = v9
            }
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                p2 = "foo bar"
                fmt.Println(p2, p2)
                fmt.Println(p1, fi11)
                _ = fi11
            }
        }
        v7 = (v7 & v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 118
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint16 = 698
    fmt.Println(v4)
    var v12 int32 = 235
    fmt.Println(v12, v12)
    switch "\n\t" {
    case "":
        if false {
            fmt.Println(v4, v12)
            var v13 int8 = -54
            _ = v13
        } else {
            var v14 int32 = (^(v12))
            _ = v14
        }
    case "\n\t":
        if false {
            if false {
                if (306 != 727) {
                    v4 += v4
                }
            } else {
                if false {
                    var v15 float32 = 0.1
                    _ = v15
                } else {
                    v12 -= 96
                    if (65535 < 30) {
                        var v16 float64 = ((float64(-2.25) - float64(((1e10 * (3.14159 * 0.1)) / 1.0))) + (1.5 - 0.0))
                        _ = v16
                    } else {
                        var v17 uint = 0
                        _ = v17
                    }
                    switch -661 {
                    case 1:
                        v4 = (v4 / (((^(v4))) | 1))
                    case 4294967295:
                        fmt.Println(v12)
                    }
                }
            }
            v18 := GenS2{F0: (895 <= genFn0(128, "", (883 > 128)))}
            _ = v18
        }
    case "foo bar":
        v19 := GenS3{F0: -1}
        _ = v19
    default:
        fmt.Println(v4)
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v4 += v4
        fmt.Println(v4)
        v21 := GenS4{F0: "\n\t"}
        _ = fi20
        _ = v21
    }
    v3[9223372036854775807] = uint32(v4)
    v22 := []int{}
    v23 := GenS2{F0: false}
    if false {
        v4 -= 255
        if true {
            var v24 int16 = 42
            v4 += (v4 * 32767)
            _ = v24
        }
        v25 := GenS3{F0: 1}
        _ = v25
    } else {
        v4 -= v4
        var v26 int64 = -707
        var v27 uint64 = 0
        _ = v26
        _ = v27
    }
    v28 := GenS5{F0: (false && true), F1: 9223372036854775807}
    var v29 uint = 70
    for fi30 := 0; fi30 < 15; fi30++ {
        fmt.Println(v12, fi30)
        _ = fi30
    }
    v29 = (^(v29))
    v29 += (v29 | (v29 | (v29 ^ (v29 - (v29 | (^(v29)))))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
