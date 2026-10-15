package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2, 190}
    var v3 int16 = 1000
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6, v3)
        fmt.Println(v3, v3)
        v7 := v4[(fi6 + -2147483648)]
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (("a" > "") || false), F1: -1, F2: -253}
    var v9 uint8 = 255
    if false {
        var v10 uint = 255
        _ = v10
    }
    v11 := GenS2{F0: 369}
    v9 -= v9
    switch "hello" {
    case "foo bar":
        v9 -= (v9 ^ (^(v9)))
    case "a":
        v3 = v3
    default:
        v12 := GenS2{F0: 146}
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        if ("foo bar" <= ((((("\n\t" + "hello") + "\n\t") + ("" + "hello")) + "foo bar") + "a")) {
            var v14 uint16 = 1
            _ = len(v4)
            v15 := v5["hello"]
            _ = v14
            _ = v15
        }
        if (!(true && (!true))) {
            v9 -= v9
            if false {
                switch -223 {
                case -1:
                    var v16 uint32 = 2147483647
                    _ = v16
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        if (-1 < 100) {
                            _ = len(v2)
                            if true {
                                var v18 int = 127
                                _ = v18
                            } else {
                                var v19 uint64 = 65
                                v3 = v3
                                _ = v19
                            }
                            v20 := GenS1{F0: (false || false), F1: -120, F2: int32(fi17)}
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            _ = len(v1)
                            var v21 float64 = -2.25
                            _ = v20
                            _ = v21
                        } else {
                            var v22 uint64 = 2
                            v22 += v22
                            _ = v22
                        }
                        _ = fi17
                    }
                default:
                    var v23 uint64 = 18446744073709551615
                    _ = v23
                }
                v3 += (v3 & -454)
                v3 = (v3 % ((-32768) | 1))
            }
        }
        v9 = v9
        _ = fi13
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v2 = append(v2, 65535)
        _ = fi24
    }
    var v25 uint64 = 4294967295
    var v26 int16 = v3
    switch 4294967295 {
    case 4294967295:
        v27 := GenS3{F0: 2, F1: 1, F2: 1}
        _ = v27
    case -634:
        var v28 uint = 42
        _ = v28
    case -245:
        switch -978 {
        case 42:
            v29 := GenS2{F0: 407}
            _ = v29
        case 2147483647:
            for fi30 := 0; fi30 < 9; fi30++ {
                v9 = (v9 ^ v9)
                fmt.Println(fi30, v9)
                v25 = 18446744073709551615
                _ = fi30
            }
        case 32768:
            v26 += -182
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
