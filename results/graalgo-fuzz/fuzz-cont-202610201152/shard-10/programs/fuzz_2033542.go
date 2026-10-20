package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float64 = -2.25
    v2 := []int{}
    var v3 int = 42
    if len(v2) > 0 {
        v2[0] = (v3 % (((v3 >> ((1) & 63))) | 1))
    }
    var v4 int32 = 0
    v5 := GenS1{F0: 1000}
    var v6 uint32 = 0
    v6 = v6
    if false {
        if true {
            var v7 float64 = float64((((0.0 * (((-2.25 - ((1e10 - -2.25) * -2.25)) / 1.5) - 0.1)) + float32(0.0)) * 3.14159))
            for fi8 := 0; fi8 < 19; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    switch "" {
                    case "foo bar":
                        var v10 float32 = 1.5
                        _ = v10
                    default:
                        fmt.Println(fi9, v1, v7)
                    }
                    _ = fi9
                }
                var v11 uint64 = 0
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: 1000}
                    _ = fi12
                    _ = v13
                }
                _ = fi8
                _ = v11
            }
            if len(v2) > 0 {
                v2[0] = -605
            }
            _ = v7
        } else {
            var v14 int64 = 1
            switch (v3 * 42) {
            case 815:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS2{F0: 0, F1: v14}
                    if false {
                        v4 = int32(v3)
                        v17 := GenS3{F0: -362, F1: (^(v6))}
                        var v18 uint64 = 18446744073709551615
                        _ = v17
                        _ = v18
                    } else {
                        switch "hello" {
                        case "":
                            if (-2.25 == float32((v1 / 1.0))) {
                                v19 := GenS2{F0: 1, F1: v14}
                                v6 = 2
                                _ = v19
                            }
                        case "hello":
                            switch "\n\t" {
                            case "hello":
                                v20 := GenS2{F0: -89, F1: v14}
                                _ = v20
                            case "\n\t":
                                var v21 float32 = float32(((float64((1.5 + 0.0)) * float64(float32(float64(0.1)))) + float64(1.5)))
                                _ = v21
                            default:
                                if false {
                                    fmt.Println(v4)
                                    for fi22 := 0; fi22 < 3; fi22++ {
                                        var v23 uint8 = 138
                                        v14 += (v14 % ((v14) | 1))
                                        _ = fi22
                                        _ = v23
                                    }
                                }
                            }
                        }
                    }
                    switch v3 {
                    case 1:
                        var v24 uint16 = 318
                        _ = v24
                    case -9223372036854775808:
                        var v25 string = "hello"
                        _ = v25
                    case 9223372036854775807:
                        fmt.Println(fi15)
                    }
                    _ = fi15
                    _ = v16
                }
            case 9223372036854775807:
                fmt.Println(v3, v6)
            default:
                _ = len(v2)
            }
            _ = v14
        }
        v3 += (v3 * (-(v3)))
    } else {
        v26 := GenS3{F0: (v4 | 0), F1: v6}
        var v27 uint16 = 0
        var v28 uint = 42
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 int64 = -9223372036854775808
    var v30 uint64 = 4294967295
    if (1 == 32767) {
        v2 = append(v2, (^(v3)))
        for fi31 := 0; fi31 < 14; fi31++ {
            fmt.Println(v3)
            _ = fi31
        }
    }
    var v32 int = 233
    if ("\n\t" >= "hello") {
        v33 := GenS4{F0: 146}
        v34 := GenS2{F0: 0, F1: -9223372036854775808}
        _ = v33
        _ = v34
    }
    v3 = -1
    if false {
        fmt.Println(v30, v1, v32)
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
