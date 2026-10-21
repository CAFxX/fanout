package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 int8, p1 uint64, p2 uint64) uint {
    var v7 string = "a"
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(p1, fi8)
        v9 := GenS1{F0: int32(p1), F1: 674, F2: -1}
        switch (fi8 / ((127) | 1)) {
        case 128:
            v10 := GenS2{F0: 137}
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    var v11 float32 = float32((float64(3.14159) * ((0.0 + (float64((((3.14159 - (1e10 / 1.0)) * 0.1) * 3.14159)) / 1e10)) * 3.14159)))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 255
}

func main() {
    var v1 bool = ((127 == 1) || false)
    var v2 float64 = 1.5
    var v3 int16 = 474
    var v4 uint = 755
    var v5 float32 = (3.14159 * (1e10 - 0.1))
    v6 := []uint32{42, 65535}
    var v12 bool = (641 == 53)
    if v12 {
        fmt.Println(v12, v3)
        if (244 == 255) {
            var v13 float64 = v2
            switch "hello" {
            case "\n\t":
                v6 = append(v6, 2147483648)
            default:
                _ = len(v6)
            }
            var v14 uint32 = 2147483647
            _ = v13
            _ = v14
        } else {
            if v1 {
                for fi15 := 0; fi15 < 8; fi15++ {
                    _ = len(v6)
                    _ = fi15
                }
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v2 = v2
                    _ = fi16
                }
            }
        }
    }
    var v17 uint32 = 2147483647
    var v18 bool = true
    switch 9223372036854775807 {
    case 127:
        v6 = append(v6, v17)
    case 175:
        switch 2147483647 {
        case 865:
            v19 := GenS3{F0: v12}
            _ = v19
        case 128:
            var v20 string = ("hello" + "\n\t")
            _ = v20
        default:
            for fi21 := 0; fi21 < 13; fi21++ {
                v5 = 1.5
                v12 = v12
                v1 = v12
                _ = fi21
            }
        }
    default:
        var v22 uint8 = 2
        _ = v22
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v23 float64 = (0.0 - v2)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v3 = (v3 * 32767)
    switch 128 {
    case 9223372036854775807:
        for fi24 := 0; fi24 < 15; fi24++ {
            v12 = v18
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v18, v5)
            case "a":
                fmt.Println(v2)
            case "hello":
                fmt.Println(v17, v3)
            default:
                v12 = v18
            }
            var v25 uint64 = 18446744073709551615
            _ = fi24
            _ = v25
        }
    case 718:
        if v1 {
            switch 899 {
            case 1:
                v26 := GenS1{F0: int32(v17), F1: int32(v4), F2: 456}
                _ = v26
            case 42:
                for fi27 := 0; fi27 < 9; fi27++ {
                    v28 := GenS1{F0: 2147483647, F1: -2147483648, F2: 2147483647}
                    v29 := GenS4{F0: 132, F1: -2147483648, F2: (^(v17))}
                    fmt.Println(v23, v17)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            v30 := GenS2{F0: 15}
            _ = v30
        } else {
            if (v12 || (2 > 65535)) {
                var v31 uint32 = v17
                _ = v31
            }
            var v32 float64 = ((v2 / ((v23 + 1e10) + (v23 - (v2 * float64(1e10))))) + float64(v5))
            fmt.Println(v12, v17, v12)
            _ = v32
        }
    case 1:
        var v33 int = 1
        _ = v33
    default:
        v34 := GenS5{F0: -9223372036854775808, F1: 178}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
