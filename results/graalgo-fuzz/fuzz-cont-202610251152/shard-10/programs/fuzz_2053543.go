package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
func genFn0(p0 uint64, p1 uint, p2 uint16) uint16 {
    v7 := []int{-387, 24, 42, 9223372036854775807}
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 string = ("a" + "")
            switch v9 {
            case "a":
                v10 := GenS1{F0: (p2 * (p2 << (((p1 + p1)) & 63))), F1: (p0 & (p0 - p0)), F2: uint64(p2)}
                _ = v10
            case "\n\t":
                v11 := GenS1{F0: 32767, F1: (p0 + (^(p0))), F2: (p0 ^ p0)}
                _ = v11
            case "foo bar":
                if false {
                    var v12 float64 = -2.25
                    _ = v12
                } else {
                    if ((p0 >> (((p1 + (p1 / ((226) | 1)))) & 63)) >= (p0 - 18446744073709551615)) {
                        fmt.Println(v9, p1)
                        v9 = v9
                    } else {
                        var v13 int = fi8
                        v14 := GenS2{F0: 0, F1: true, F2: (0 > 129)}
                        _ = v13
                        _ = v14
                    }
                    v9 = (v9 + "\n\t")
                    for fi15 := 0; fi15 < 13; fi15++ {
                        p0 = 42
                        _ = fi15
                    }
                }
            }
            _ = fi8
            _ = v9
        }
    case "a":
        v7 = append(v7, 943)
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = 1000
        p0 -= p0
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 903, F1: ((1 > -1) || false), F2: false}
    var v19 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    v2 := []uint32{2147483647}
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    v5 := []int{220, 9223372036854775807}
    var v6 int8 = 111
    for fi20 := 0; fi20 < 8; fi20++ {
        switch (fi20 % ((fi20) | 1)) {
        case 32768:
            v3 += (v3 << (((v3 << ((2) & 63))) & 63))
        case 921:
            v6 = (^(v6))
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        switch fi20 {
        case -2147483648:
            v3 -= 255
        case -242:
            for fi21 := 0; fi21 < 14; fi21++ {
                fmt.Println(v3, fi21, fi21)
                _ = fi21
            }
        case 2147483647:
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 uint32 = uint32(v3)
                _ = fi22
                _ = v23
            }
        default:
            if false {
                if true {
                    fmt.Println(v3)
                }
                if ((fi20 & (fi20 * fi20)) == (fi20 << (((v3 - (v3 * (v3 & 470)))) & 63))) {
                    fmt.Println(fi20, v3)
                }
                v6 -= v6
            }
        }
        _ = fi20
    }
    _ = len(v4)
    var v24 bool = true
    v24 = v24
    fmt.Println(v3, v3)
    v24 = v24
    var v25 uint16 = 0
    var v26 float64 = ((((1.5 / 1.0) + 3.14159) / float64(1.5)) / (float64(((-2.25 / 0.1) / (0.1 / 1e10))) * (0.1 * (-2.25 + 0.1))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
