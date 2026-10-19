package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 uint32, p1 uint32, p2 uint16) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int = fi5
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = -97
            _ = fi7
            _ = v8
        }
        var v9 int16 = -775
        _ = fi5
        _ = v6
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 649, F2: -527}
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 128
}
func genFn1(p0 bool, p1 int64, p2 int32) uint64 {
    v11 := []int{47, -2147483648, 2147483647}
    _ = len(v11)
    if p0 {
        var v12 float64 = 0.0
        var v13 uint16 = 2
        switch "a" {
        case "foo bar":
            var v14 float32 = (1e10 - 1e10)
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 7; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    fmt.Println(fi16, fi15)
                    switch (-(fi16)) {
                    case 9223372036854775807:
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case 1:
                        if ((1e10 / 0.1) > 0.0) {
                            var v17 float64 = v12
                            var v18 uint8 = 42
                            _ = v17
                            _ = v18
                        } else {
                            var v19 bool = p0
                            _ = v19
                        }
                    case -2147483648:
                        var v20 uint32 = 426
                        _ = v20
                    }
                    v13 += v13
                    _ = fi16
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi15
            }
        case "\n\t":
            if len(v11) > 0 {
                v11[0] = 127
            }
        default:
            v11 = append(v11, -68)
        }
        _ = v12
        _ = v13
    } else {
        v21 := GenS2{F0: (42 <= 65535), F1: 2147483648}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 752
}

func main() {
    var v1 uint8 = 0
    v2 := []string{}
    var v3 bool = true
    var v4 float32 = (-2.25 + (1.5 * 3.14159))
    for fi22 := 0; fi22 < 14; fi22++ {
        v3 = (!("foo bar" <= ("hello" + "foo bar")))
        v4 = -2.25
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS3{F0: 778, F1: int16(v1), F2: (fi24 | fi24)}
        _ = fi24
        _ = v25
    }
    fmt.Println(v3, v1)
    var v26 float64 = float64((v4 + -2.25))
    var v27 float64 = 0.0
    switch "\n\t" {
    case "a":
        fmt.Println(v26)
    default:
        fmt.Println(v26, v26, v4)
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS3{F0: -32768, F1: 1, F2: (fi28 % (((fi28 >> ((42) & 63))) | 1))}
        _ = fi28
        _ = v29
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
