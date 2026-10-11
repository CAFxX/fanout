package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16) uint32 {
    v7 := []int{128, 9223372036854775807, -9223372036854775808, 42}
    p0 = 255
    fmt.Println(p0)
    _ = p0
    _ = v7
    return uint32(p0)
}
func genFn1(p0 int8, p1 bool, p2 uint16) int16 {
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (fi8 | (fi8 >> ((3) & 63)))}
        if (!true) {
            for fi10 := 0; fi10 < 12; fi10++ {
                p0 = int8(p2)
                if (1 > (p2 ^ p2)) {
                    if p1 {
                        v11 := GenS2{F0: 1, F1: -559}
                        p2 = 2
                        p0 = (p0 | (p0 - p0))
                        _ = v11
                    } else {
                        v12 := GenS2{F0: 206, F1: -427}
                        _ = v12
                    }
                }
                _ = fi10
            }
            var v13 int32 = 1
            _ = v13
        }
        fmt.Println(p1, p2, p0)
        _ = fi8
        _ = v9
    }
    switch -9223372036854775808 {
    case 4294967295:
        p0 = p0
    default:
        p2 += 144
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 bool = (5 >= 706)
    var v2 uint16 = 801
    var v3 uint16 = v2
    var v4 uint16 = v2
    var v5 uint = 255
    v6 := make(map[string]uint32)
    if v1 {
        v4 = (v2 | (^(v4)))
        var v14 int64 = -950
        _ = len(v6)
        _ = v14
    }
    var v15 int = -9223372036854775808
    v5 += (^(v5))
    for fi16 := 0; fi16 < 17; fi16++ {
        switch "\n\t" {
        case "foo bar":
            if v1 {
                fmt.Println(v15)
            } else {
                var v17 uint16 = v4
                var v18 uint16 = (v4 / ((v2) | 1))
                switch "foo bar" {
                case "":
                    switch "" {
                    case "a":
                        for fi19 := 0; fi19 < 9; fi19++ {
                            v1 = true
                            _ = fi19
                        }
                    default:
                        v17 = v2
                    }
                }
                _ = v17
                _ = v18
            }
        default:
            var v20 float32 = (1e10 - float32(float64(3.14159)))
            _ = v20
        }
        switch (v15 >> ((v5) & 63)) {
        case -165:
            fmt.Println(v4, v2, v1)
        case -2147483648:
            if v1 {
                var v21 int8 = 42
                _ = v21
            }
        default:
            v3 -= 42
        }
        _ = fi16
    }
    var v22 string = ("hello" + "a")
    v3 = 42
    v22 = (v22 + v22)
    v5 += 42
    if true {
        v6[v22] = 0
    }
    v4 -= (^(v2))
    switch v22 {
    case "\n\t":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := []int{v15, v15, (v15 + -117), (^(fi23))}
            v5 += v5
            _ = fi23
            _ = v24
        }
    case "foo bar":
        v25 := []int{}
        _ = v25
    case "":
        switch v22 {
        case "a":
            v26 := []int{}
            _ = v26
        case "foo bar":
            v3 -= (^(v2))
        case "":
            if v1 {
                v15 = (v15 & v15)
                var v27 int64 = -9223372036854775808
                v15 -= v15
                _ = v27
            }
        default:
            _ = len(v6)
        }
    default:
        fmt.Println(v2, v4)
    }
    var v28 uint8 = 42
    switch (v15 << ((v5) & 63)) {
    case 9223372036854775807:
        v22 = v22
    case 127:
        v29 := GenS2{F0: v5, F1: 596}
        _ = v29
    case 2147483647:
        if (uint32(v5) < 908) {
            if v1 {
                v30 := GenS1{F0: (-(v15))}
                _ = v30
            }
            v31 := GenS3{F0: -878}
            for fi32 := 0; fi32 < 16; fi32++ {
                fmt.Println(v28, v22, v4)
                v33 := GenS3{F0: -826}
                _ = fi32
                _ = v33
            }
            _ = v31
        } else {
            var v34 int64 = 861
            v28 = v28
            var v35 uint16 = v2
            _ = v34
            _ = v35
        }
    }
    v36 := GenS3{F0: -857}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
