package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{845}
    v3 := make(map[uint32]string)
    var v4 float64 = (0.0 / -2.25)
    v5 := GenS1{F0: 1, F1: v1}
    v4 = (v4 * v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 uint = 1
    switch -9223372036854775808 {
    case 509:
        for fi7 := 0; fi7 < 12; fi7++ {
            if (!(763 >= 44)) {
                fmt.Println(v4, v6, v6)
            }
            switch -464 {
            case -2147483648:
                v1 = v1
            default:
                if false {
                    v8 := GenS1{F0: 113, F1: (^(v1))}
                    _ = v8
                }
            }
            var v9 int8 = -112
            _ = fi7
            _ = v9
        }
    default:
        fmt.Println(v1, v6, v4)
    }
    if (837 < 1) {
        var v10 int32 = -1
        var v11 int32 = (v10 | 42)
        v12 := GenS2{F0: 65535, F1: uint(v1)}
        _ = v10
        _ = v11
        _ = v12
    }
    v6 += uint(v1)
    var v13 uint8 = 128
    v14 := GenS2{F0: 65535, F1: (v6 - (^(v6)))}
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            v4 = v4
            v2 = append(v2, 2147483648)
            _ = fi16
        }
        var v17 uint32 = 700
        switch ("hello" + ("" + "foo bar")) {
        case "":
            v13 += 137
        case "\n\t":
            var v18 uint32 = (^(v17))
            _ = v18
        }
        _ = fi15
        _ = v17
    }
    fmt.Println(v4, v4)
    var v19 int8 = -110
    if false {
        v2 = append(v2, 734)
        switch ("a" + "a") {
        case "foo bar":
            v13 -= v13
        case "a":
            fmt.Println(v19)
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 int = fi20
            fmt.Println(v4, v21, v19)
            _ = fi20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 9; fi22++ {
            switch fi22 {
            case -517:
                v13 -= 144
            case 266:
                for fi23 := 0; fi23 < 10; fi23++ {
                    _ = len(v2)
                    v24 := GenS1{F0: 42, F1: v1}
                    switch "\n\t" {
                    case "\n\t":
                        v25 := GenS1{F0: v19, F1: v1}
                        _ = v25
                    case "a":
                        var v26 uint64 = 18446744073709551615
                        v27 := GenS2{F0: 42, F1: v6}
                        v28 := GenS2{F0: 2, F1: v6}
                        _ = v26
                        _ = v27
                        _ = v28
                    case "hello":
                        v13 = (v13 / (((v13 / ((127) | 1))) | 1))
                    }
                    var v29 uint64 = 42
                    _ = fi23
                    _ = v24
                    _ = v29
                }
            case -383:
                switch "a" {
                case "hello":
                    v30 := GenS2{F0: 2, F1: v6}
                    _ = v30
                }
            }
            delete(v3, 576)
            if true {
                v6 -= (v6 << ((v6) & 63))
            }
            _ = fi22
        }
        v1 += (v1 | (v1 / ((115) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
