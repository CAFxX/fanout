package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int16; F2 int32 }

func main() {
    var v1 float64 = float64(0.0)
    v2 := []int{}
    var v3 uint16 = 642
    var v4 uint8 = 2
    v5 := make(map[uint32]bool)
    var v6 int = -2147483648
    if false {
        v3 += v3
        if (2 == (v4 & (v4 & v4))) {
            delete(v5, 353)
            var v7 uint32 = 374
            _ = v7
        }
    }
    v8 := GenS1{F0: -1}
    v9 := GenS2{F0: (^(v3)), F1: (false || (1 < uint64(v4)))}
    v10 := GenS1{F0: 32767}
    if true {
        var v11 int32 = 0
        _ = v11
    }
    _ = len(v5)
    switch (v6 % ((9223372036854775807) | 1)) {
    case 127:
        fmt.Println(v4)
    case -1:
        if (!false) {
            v1 -= v1
            delete(v5, uint32(v3))
        }
    case 4294967295:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v6, v4)
        case "hello":
            v12 := GenS2{F0: (v3 << ((509) & 63)), F1: true}
            _ = v12
        }
    default:
        if (v1 < ((3.14159 - v1) + v1)) {
            var v13 int64 = -1
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS2{F0: (v3 / (((v3 + v3)) | 1)), F1: (!(!(42 > 2)))}
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint8 = (v4 | 255)
                    v4 = v17
                    v4 = (v17 << ((111) & 63))
                    _ = fi16
                    _ = v17
                }
                v6 = 752
                _ = fi14
                _ = v15
            }
        }
    }
    if (!true) {
        fmt.Println(v1, v1)
    }
    v18 := v5[2147483647]
    v4 += v4
    var v19 int8 = -128
    if (v6 < (v6 | (v6 - 1))) {
        if ((float32((1e10 * 3.14159)) + (0.0 / (((((float32(v1) - ((((0.1 * 0.1) - (1.5 + float32(v1))) - 3.14159) * -2.25)) / 1.0) - ((1e10 - 0.1) - -2.25)) - (float32((v1 + v1)) + (3.14159 + 1e10))) + 1e10))) >= ((1.5 / 1.0) + float32(((v1 * ((((v1 + 1.5) + v1) - v1) + v1)) + v1)))) {
            v2 = append(v2, -165)
            var v20 bool = (((-592 < -484) && v18) || (1000 < 52))
            _ = v20
        }
    } else {
        if (0.1 > v1) {
            v19 = 127
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint16 = v3
                    v24 := GenS3{F0: "a"}
                    if (((v19 % (((^(v19))) | 1)) < (v19 << ((0) & 63))) && v18) {
                        var v25 uint64 = 0
                        var v26 int64 = 0
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                var v27 uint16 = (^(v3))
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 bool = (!false)
                    v29 = (-1 != int16(v3))
                    v19 += (v19 * v19)
                    if (996 >= 617) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v29)
                        var v30 int16 = 0
                        _ = v30
                    } else {
                        delete(v5, 612)
                        var v31 int64 = 1
                        var v32 int32 = -944
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                }
                _ = fi21
                _ = v27
            }
            v19 += v19
        }
        v19 = -41
    }
    v5[4294967295] = (32767 < v3)
    v33 := v5[2147483648]
    switch "" {
    case "hello":
        v34 := GenS4{F0: 506, F1: -1, F2: -642}
        _ = v34
    case "a":
        fmt.Println(v6, v18)
    case "":
        switch "foo bar" {
        case "foo bar":
            var v35 uint = 42
            _ = v35
        case "\n\t":
            var v36 int = v6
            _ = v36
        case "":
            v37 := GenS3{F0: "a"}
            _ = v37
        default:
            fmt.Println(v33, v6)
        }
    default:
        if (233 >= (v4 ^ (v4 / ((1) | 1)))) {
            for fi38 := 0; fi38 < 11; fi38++ {
                v39 := GenS1{F0: -826}
                var v40 int = v6
                var v41 int64 = 1
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        } else {
            var v42 uint32 = 113
            for fi43 := 0; fi43 < 13; fi43++ {
                v42 += (v42 & v42)
                _ = fi43
            }
            _ = v42
        }
    }
    switch v6 {
    case 32768:
        v5[0] = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
