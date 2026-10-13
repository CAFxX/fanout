package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 bool; F2 bool }

func main() {
    v1 := []uint32{0}
    v2 := make(map[uint32]uint32)
    var v3 bool = (65535 < 0)
    var v4 bool = false
    v5 := make(map[int]bool)
    var v6 float64 = 0.1
    switch "a" {
    case "foo bar":
        fmt.Println(v3, v6)
    default:
        var v7 float32 = ((1.5 / 1.0) * 0.0)
        _ = v7
    }
    switch "" {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        v6 += -2.25
    case "":
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42, F1: 0}
            _ = len(v2)
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: 9223372036854775807, F1: 0}
                var v12 int64 = 42
                v4 = (42 < 0)
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    }
    var v13 uint32 = 1
    v13 += v13
    v6 += v6
    for fi14 := 0; fi14 < 19; fi14++ {
        v2[v13] = v13
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS2{F0: 32767, F1: 2}
            _ = fi15
            _ = v16
        }
        var v17 int16 = 42
        _ = fi14
        _ = v17
    }
    switch -9223372036854775808 {
    case -883:
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v3, v6)
        case 2147483647:
            v18 := GenS2{F0: 0, F1: 521}
            _ = v18
        default:
            v4 = (!((32767 >= 32767) && true))
        }
    case 471:
        for fi19 := 0; fi19 < 2; fi19++ {
            switch "foo bar" {
            case "a":
                v3 = false
            default:
                switch "\n\t" {
                case "":
                    fmt.Println(fi19, v3, fi19)
                case "hello":
                    var v20 int32 = 2147483647
                    _ = v20
                }
            }
            var v21 float32 = 1.5
            var v22 int8 = -56
            _ = fi19
            _ = v21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if v3 {
            var v24 int = fi23
            _ = v24
        } else {
            fmt.Println(v4, v4)
            delete(v5, (fi23 / ((fi23) | 1)))
        }
        _ = fi23
    }
    if v3 {
        var v25 int32 = 42
        for fi26 := 0; fi26 < 13; fi26++ {
            _ = len(v1)
            v25 -= 1
            _ = fi26
        }
        _ = v25
    } else {
        var v27 string = ""
        var v28 string = "foo bar"
        var v29 uint = 115
        _ = v27
        _ = v28
        _ = v29
    }
    if (119 == 42) {
        v3 = (!(v4 && (2 != 658)))
    }
    v30 := GenS3{F0: -1, F1: (!v3), F2: (392 > 341)}
    v13 = v13
    for fi31 := 0; fi31 < 4; fi31++ {
        v1 = append(v1, v13)
        v13 = 2147483648
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
