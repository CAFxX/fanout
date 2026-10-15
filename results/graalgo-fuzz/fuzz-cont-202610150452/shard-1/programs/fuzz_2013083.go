package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 bool; F1 int8; F2 bool }
func genFn0(p0 int8, p1 int16, p2 uint8) int16 {
    var v7 float64 = 0.1
    v8 := []int{2147483647}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []uint32{42, 0, 65535, 42}
    var v2 uint16 = 849
    var v3 int32 = 42
    v4 := []int{4294967295, -2147483648, 32767, 42}
    v5 := []uint32{264, 259}
    var v6 uint8 = 42
    switch 2147483647 {
    case 128:
        v3 -= 449
    }
    if false {
        fmt.Println(v6, v2)
    }
    fmt.Println(v6, v2)
    for fi9 := 0; fi9 < 4; fi9++ {
        v6 -= (v6 % (((v6 ^ 68)) | 1))
        _ = fi9
    }
    var v10 float64 = 0.0
    if (42 > 34) {
        v10 = (1.5 * v10)
    } else {
        fmt.Println(v10, v6)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v11 := GenS1{F0: 732, F1: 42, F2: false}
    v3 = 635
    var v12 float32 = (float32((v10 / (v10 / (((v10 + v10) + float64(0.0)) / float64(0.1))))) * (((0.0 + ((3.14159 + 1e10) - (0.1 + 0.1))) * 1.5) + (-2.25 + (0.1 + float32(3.14159)))))
    v10 -= float64((1e10 * 1e10))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v13 uint = 2
        for fi14 := 0; fi14 < 7; fi14++ {
            v2 = v2
            _ = len(v5)
            fmt.Println(v13, v2, v2)
            _ = fi14
        }
        v3 -= v3
        _ = v13
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: genFn0(100, 42, v6), F1: 100, F2: false}
        v17 := GenS2{F0: (v3 / (((v3 % (((v3 & (-(v3)))) | 1))) | 1)), F1: "hello", F2: 0}
        var v18 uint = 255
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if true {
        v12 += v12
        switch 32768 {
        case 1:
            v19 := GenS2{F0: -2147483648, F1: "foo bar", F2: 9223372036854775807}
            _ = v19
        case 32767:
            switch 912 {
            case 0:
                v20 := GenS3{F0: (false && (true || (0 <= 18446744073709551615))), F1: 127, F2: true}
                _ = v20
            case 2147483647:
                fmt.Println(v12, v6, v2)
            default:
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 int32 = v3
                    _ = fi21
                    _ = v22
                }
            }
        case -991:
            var v23 int8 = 0
            _ = v23
        }
        if (483 < (v2 & v2)) {
            v24 := GenS3{F0: true, F1: 127, F2: false}
            switch "hello" {
            case "foo bar":
                if true {
                    v25 := GenS3{F0: false, F1: 1, F2: false}
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS1{F0: 0, F1: 64, F2: (v2 < 0)}
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 float32 = (v12 - v12)
                        _ = fi28
                        _ = v29
                    }
                }
            case "\n\t":
                v3 -= v3
            case "a":
                var v30 uint32 = 42
                _ = v30
            }
            _ = v24
        }
    }
    v31 := GenS2{F0: (v3 << ((0) & 63)), F1: "hello", F2: 9223372036854775807}
    var v32 int32 = 798
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
