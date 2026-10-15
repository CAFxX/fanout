package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint8; F1 int }

func main() {
    var v1 uint64 = 46
    v2 := []uint32{65535}
    v3 := []int{4294967295, -735, 4294967295, 534}
    v4 := []uint32{4294967295, 2147483648, 2147483647}
    var v5 int8 = -92
    var v6 uint64 = 42
    fmt.Println(v5, v6)
    v1 = 18446744073709551615
    switch 237 {
    case 872:
        for fi7 := 0; fi7 < 20; fi7++ {
            v5 += v5
            _ = fi7
        }
    case 69:
        for fi8 := 0; fi8 < 15; fi8++ {
            v6 -= 215
            for fi9 := 0; fi9 < 1; fi9++ {
                if (!(!(255 != 454))) {
                    v6 += v6
                    v10 := GenS1{F0: 18446744073709551615, F1: -324}
                    var v11 int64 = 465
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            _ = fi8
        }
    case -1:
        var v12 float64 = 3.14159
        _ = v12
    default:
        var v13 float32 = 1.5
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                switch fi15 {
                case 9223372036854775807:
                    var v16 uint8 = 42
                    _ = v16
                case -2147483648:
                    v3 = append(v3, (-(fi15)))
                case -33:
                    v5 = v5
                }
                var v17 int = 2147483647
                var v18 uint16 = 468
                _ = fi15
                _ = v17
                _ = v18
            }
            fmt.Println(v6)
            _ = len(v4)
            _ = fi14
        }
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS2{F0: ("" + "a"), F1: 409}
        _ = fi19
        _ = v20
    }
    fmt.Println(v6)
    for fi21 := 0; fi21 < 11; fi21++ {
        switch 0 {
        case 4294967295:
            var v22 uint16 = 32767
            _ = v22
        case -9223372036854775808:
            v23 := GenS3{F0: 1, F1: v6, F2: fi21}
            _ = v23
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            v24 := GenS2{F0: ("a" + "hello"), F1: -9223372036854775808}
            _ = v24
        } else {
            v25 := GenS4{F0: 232, F1: v5}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v25
        }
        _ = fi21
    }
    v5 = v5
    var v26 uint8 = 127
    v6 = v6
    v27 := GenS2{F0: "foo bar", F1: -557}
    v28 := GenS5{F0: v26, F1: 565}
    var v29 string = ""
    v30 := GenS4{F0: (v26 | 2), F1: -128}
    for fi31 := 0; fi31 < 1; fi31++ {
        switch 513 {
        case -9223372036854775808:
            if (!true) {
                fmt.Println(fi31, v5, v29)
            } else {
                v1 -= v6
                var v32 uint32 = 65535
                _ = v32
            }
        case 127:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32767:
            if (42 < (fi31 % ((-9223372036854775808) | 1))) {
                v5 -= 1
                var v33 uint32 = 311
                v34 := GenS4{F0: (v26 >> ((uint(v26)) & 63)), F1: (^(v5))}
                _ = v33
                _ = v34
            }
        default:
            v29 = v29
        }
        v26 += v26
        var v35 float32 = 3.14159
        _ = fi31
        _ = v35
    }
    fmt.Println(v5, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
