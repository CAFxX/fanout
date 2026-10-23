package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint8, p1 int16, p2 int16) bool {
    v7 := []int{}
    fmt.Println(p2, p0, p0)
    v8 := GenS1{F0: 65535, F1: false, F2: -59}
    v7 = append(v7, 128)
    if len(v7) > 0 {
        v7[0] = -253
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (2 <= 42)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    v3 := make(map[string]uint32)
    var v4 int64 = -1
    v5 := make(map[uint32]int)
    var v6 uint = 190
    fmt.Println(v4, v6)
    switch int(v4) {
    case 32767:
        var v9 int64 = -668
        _ = v9
    case 77:
        var v10 int64 = (-(v4))
        _ = v10
    case 315:
        v6 += v6
    }
    v11 := GenS2{F0: (62 >= 42), F1: (v6 | 0), F2: v2}
    if (2 == 74) {
        v12 := []int{159}
        _ = v12
    }
    v2 = v2
    v13 := []int{128, 2147483647}
    if v2 {
        v14 := GenS2{F0: false, F1: (v6 % ((255) | 1)), F2: v2}
        if (-128 == 108) {
            for fi15 := 0; fi15 < 8; fi15++ {
                v6 -= (v6 << ((v6) & 63))
                _ = fi15
            }
        } else {
            if (false && v2) {
                var v16 int64 = (v4 + -1)
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    v18 := GenS1{F0: 65535, F1: (1000 >= 1), F2: 1}
                    _ = fi17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                v6 = (v6 << (((v6 - v6)) & 63))
                switch "" {
                case "hello":
                    v20 := GenS2{F0: v2, F1: v6, F2: v2}
                    _ = v20
                }
                _ = fi19
            }
            var v21 int16 = 1
            _ = v21
        }
        v22 := GenS3{F0: false, F1: 755}
        _ = v14
        _ = v22
    } else {
        switch ("" + ("hello" + "")) {
        case "\n\t":
            v2 = v2
        case "hello":
            v4 -= (v4 / (((v4 & (v4 | (v4 & v4)))) | 1))
        default:
            var v23 uint = (v6 | (v6 - 884))
            _ = v23
        }
        fmt.Println(v4, v2)
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    v24 := GenS1{F0: 2, F1: true, F2: 86}
    if (225 > 0) {
        var v25 int16 = -591
        _ = v25
    }
    fmt.Println(v2, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
